package scraper

import (
	"house4sale/models"
	"log"
)

func Aggregate() ([]models.House, error) {

    log.Println("Calling Aggregate()")

    fastighetsbyran, _ := ScrapeFastighetsbyran()
    skandiamaklarna, _ := ScrapeSkandiaMaklarna()
    lansforsakringar, err := ScrapeLansforsakringar()
    if err != nil {
        return nil, err
    }

    all := []models.House{}
    all = append(all, fastighetsbyran...)
    all = append(all, skandiamaklarna...)
    all = append(all, lansforsakringar...)

    return all, nil
}
