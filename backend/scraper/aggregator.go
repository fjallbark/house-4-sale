package scraper

import (
	"house4sale/models"
	"log"
)

func Aggregate() ([]models.House, error) {

    log.Println("Calling Aggregate()")

    fastighetsbyran, _ := ScrapeFastighetsbyran()
    skandiamaklarna, _ := ScrapeSkandiaMaklarna()
    lansforsakringar, _ := ScrapeLansforsakringar()
    bjurfors, err := ScrapeBjurfors()
    if err != nil {
        return nil, err
    }

    all := []models.House{}
    all = append(all, fastighetsbyran...)
    all = append(all, skandiamaklarna...)
    all = append(all, lansforsakringar...)
    all = append(all, bjurfors...)

    return all, nil
}
