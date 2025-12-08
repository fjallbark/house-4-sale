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
    bjurfors, _ := ScrapeBjurfors()
    svenskFast, err := ScrapeSvenskFast()

    if err != nil {
        return nil, err
    }

    all := []models.House{}
    all = append(all, fastighetsbyran...)
    all = append(all, skandiamaklarna...)
    all = append(all, lansforsakringar...)
    all = append(all, bjurfors...)
    all = append(all, svenskFast...)

    return all, nil
}
