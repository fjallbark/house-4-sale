package scraper

import (
	"house4sale/models"
	"log"
)

func Aggregate() ([]models.House, error) {

    log.Println("Calling Aggregate()")

    fastighetsbyran, err := ScrapeFastighetsbyran()
    if err != nil {
        return nil, err
    }

    // Lägg till fler siter här:
    // booli, _ := ScrapeBooli()
    // boneo, _ := ScrapeBoneo()

    all := []models.House{}
    all = append(all, fastighetsbyran...)
    // all = append(all, booli...)
    // all = append(all, boneo...)

    return all, nil
}
