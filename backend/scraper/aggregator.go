package scraper

import (
	"house4sale/models"
	"log"
)

func Aggregate() ([]models.House, error) {

    log.Println("Calling Aggregate()")

    fastighetsbyran, _ := ScrapeFastighetsbyran()
    skandiamaklarna, err := ScrapeSkandiaMaklarna()
    if err != nil {
        return nil, err
    }

    all := []models.House{}
    all = append(all, fastighetsbyran...)
    all = append(all, skandiamaklarna...)

    return all, nil
}
