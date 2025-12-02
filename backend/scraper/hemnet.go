package scraper

import (
    "net/http"
    "strings"

    "github.com/PuerkitoBio/goquery"
    "house4sale/models"

    "log"
)

func ScrapeHemnet() ([]models.House, error) {
    log.Println("Calling ScrapeHemnet()")

    resp, err := http.Get("https://www.fastighetsbyran.com/sv/sverige/till-salu")

    log.Println("Resp", resp)

    if err != nil {
        log.Println("Err is nil")
        return nil, err
    }
    defer resp.Body.Close()

    log.Println("Resp body", resp.Body)

    doc, err := goquery.NewDocumentFromReader(resp.Body)

    log.Println("Doc", doc)

    if err != nil {
        log.Println("Err is nil")
        return nil, err
    }

    var results []models.House

    log.Println("Results", results)

    doc.Find(".listing-card").Each(func(i int, s *goquery.Selection) {

        log.Println("Find")

        title := strings.TrimSpace(s.Find(".listing-card__title").Text())
        price := strings.TrimSpace(s.Find(".listing-card__price").Text())
        address := strings.TrimSpace(s.Find(".listing-card__location").Text())
        image, _ := s.Find("img").Attr("src")
        url, _ := s.Find("a").Attr("href")

        if url != "" && !strings.HasPrefix(url, "http") {
            url = "https://www.hemnet.se" + url
        }

        results = append(results, models.House{
            Title:   title,
            Price:   price,
            Address: address,
            Image:   image,
            Url:     url,
            Source:  "hemnet",
        })
    })

    return results, nil
}
