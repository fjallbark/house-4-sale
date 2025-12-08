package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"house4sale/models"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeSvenskFast() ([]models.House, error) {
    log.Println("ScrapeSvenskFast")

    url := "https://www.svenskfast.se/Sweden/sv/search/object?l=skane%2Fhoor&t=Villa"

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("non-200: %s", resp.Status)
    }

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        return nil, err
    }

    var houses []models.House

    doc.Find("script[type='application/ld+json']").Each(func(i int, s *goquery.Selection) {
        jsonText := s.Text()
        var data map[string]interface{}
        err := json.Unmarshal([]byte(jsonText), &data)
        if err != nil {
            return
        }

        // Kontrollera att det är en RealEstateListing
        if data["@type"] != "RealEstateListing" {
            return
        }

        title, _ := data["name"].(string)
        url, _ := data["url"].(string)

        addressMap, _ := data["address"].(map[string]interface{})
        street, _ := addressMap["streetAddress"].(string)
        locality, _ := addressMap["addressLocality"].(string)
        address := locality + ", " + street

        priceMap, _ := data["price"].(map[string]interface{})
        priceStr := ""
        if priceMap != nil {
            priceVal, _ := priceMap["value"].(string)
            priceStr = priceVal + " kr"
        }

        rooms := ""
        if v, ok := data["numberOfRooms"].(string); ok {
            rooms = v + " rok"
        }

        area := ""
        if floorSize, ok := data["floorSize"].(map[string]interface{}); ok {
            if v, ok := floorSize["value"].(float64); ok {
                area = fmt.Sprintf("%.0f kvm", v)
            }
        }

        images := []string{}
        if imgList, ok := data["image"].([]interface{}); ok {
            for _, img := range imgList {
                if imgStr, ok := img.(string); ok {
                    images = append(images, imgStr)
                }
            }
        }

        img := ""
        if len(images) > 0 {
            img = images[0]
        }

        houses = append(houses, models.House{
            Title:        title,
            Address:      address,
            Price:        priceStr,
            Rooms:        rooms,
            SquareMeters: area,
            Image:        img,
            Url:          url,
            Source:       "SvenskFast",
        })
    })

    log.Println("houses svenskfast", len(houses))
    return houses, nil
}
