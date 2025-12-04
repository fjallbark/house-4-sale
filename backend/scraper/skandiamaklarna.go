package scraper

import (
	"fmt"
	"log"
	"net/http"

	"house4sale/models"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeSkandiaMaklarna() ([]models.House, error) {
    log.Println("ScrapeSkandiaMaklarna")

    url := "https://www.skandiamaklarna.se/hitta-hem/Villa/H%C3%B6%C3%B6r?layout=Default&municipalityId=90&municipalityId-90-label=H%C3%B6%C3%B6r+%28Kommun%29&typeId=2%7C61"

    res, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    if res.StatusCode != 200 {
        return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
    }

    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        return nil, err
    }

    var houses []models.House

    doc.Find(".estate-search-result-item").Each(func(i int, s *goquery.Selection) {
    url, _ := s.Find("a").Attr("href")

    image, _ := s.Find("img").Attr("src")

    title := s.Find("h3").Text()
    neighborhood := s.Find("hgroup p").Text()
    address := fmt.Sprintf("%s, %s", title, neighborhood)

    rooms := s.Find(".quick-facts span.value").Eq(0).Text()
    squareMeters := s.Find(".quick-facts span.value").Eq(1).Text()
    price := s.Find(".quick-facts span").Eq(2).Text()

    houses = append(houses, models.House{
        Title:        title,
        Price:        price,
        Rooms:        rooms,
        SquareMeters: squareMeters,
        Address:      address,
        Image:        image,
        Url:          url,
        Source:       "SkandiaMäklarna",
    })
})


    log.Println("houses skandiamaklarna", len(houses))

    return houses, nil
}

