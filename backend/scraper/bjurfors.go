package scraper

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"house4sale/models"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeBjurfors() ([]models.House, error) {
    log.Println("ScrapeBjurfors")

    url := "https://www.bjurfors.se/sv/tillsalu/skane/hoor/hoor/?type=Bungalow&pmax=6000000&rcmin=5"

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

    doc.Find(".c-object-card").Each(func(i int, s *goquery.Selection) {
        title := strings.TrimSpace(s.Find(".c-object-card__heading").Text())
        link, _ := s.Find(".c-object-card__link").Attr("href")
        if !strings.HasPrefix(link, "http") {
            link = "https://www.bjurfors.se" + link
        }

        cityArea := strings.TrimSpace(s.Find(".c-object-card__city-area").Text())
        address := strings.TrimSpace(s.Find(".c-object-card__address").Text())

        var rooms, area, price string
        s.Find(".c-object-card__meta").Each(func(i int, m *goquery.Selection) {
            text := strings.TrimSpace(m.Text())
            if strings.Contains(text, "rum") {
                rooms = text
            } else if strings.Contains(text, "kvm") {
                area = text
            } else if strings.Contains(text, "kr") {
                price = text
            }
        })

        img, _ := s.Find("img.c-object-card__image").Attr("src")
        if strings.HasPrefix(img, "/") {
            img = "https://www.bjurfors.se" + img
        }

        houses = append(houses, models.House{
            Title:        title,
            Address:      cityArea + ", " + address,
            Price:        price,
            Rooms:        rooms,
            SquareMeters: area,
            Image:        img,
            Url:          link,
            Source:       "Bjurfors",
        })
    })

    log.Println("houses bjurfors", len(houses))
    return houses, nil
}
