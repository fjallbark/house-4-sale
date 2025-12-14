package scraper

import (
	"house4sale/models"
	"log"
)

func Aggregate() ([]models.House, error) {

    log.Println("Calling Aggregate()")

    /*fastighetsbyran, _ := ScrapeFastighetsbyran()
    skandiamaklarna, _ := ScrapeSkandiaMaklarna()
    lansforsakringar, _ := ScrapeLansforsakringar()
    bjurfors, _ := ScrapeBjurfors()
    svenskFast, err := ScrapeSvenskFast()*/
    erikOlsson, err := ScrapeErikOlsson()

    if err != nil {
        return nil, err
    }

    all := []models.House{}
    /*all = append(all, fastighetsbyran...)
    all = append(all, skandiamaklarna...)
    all = append(all, lansforsakringar...)
    all = append(all, bjurfors...)
    all = append(all, svenskFast...)*/
    all = append(all, erikOlsson...)

    return all, nil
}
