package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"house4sale/models"
)

type fastAPIResponseLans struct {
    TotalLength int          `json:"totalLength"`
    Estates     []EstateItem `json:"estates"`
}

type EstateItem struct {
    ID            string  `json:"id"`
    Url           string  `json:"url"`
    StreetAddress string  `json:"streetAddress"`
    StartPrice    float64 `json:"startPrice"`
    Livingspace   float64 `json:"livingspace"`
    NumberOfRooms float64 `json:"numberOfRooms"`
    HeaderImage   string  `json:"headerImage"`
    PublishDate   string  `json:"publishDate"`
    Area          string  `json:"area"`
}

func ScrapeLansforsakringar() ([]models.House, error) {
    log.Println("ScrapeLansforsakringar")

    url := "https://app-lansfast-api.azurewebsites.net/api/Estates/GetForFilter" +
        "?page=1" +
        "&municipality=H%C3%B6%C3%B6r" +
        "&estateType=Villa" +
        "&priceMax=6500000" +
        "&roomsMin=5" +
        "&roomsMax=10" +
        "&livingSpaceMin=140" +
        "&livingSpaceMax=200" +
        "&plotAreaMin=500" +
        "&plotAreaMax=15000" +
        "&showSold=false" +
        "&sortOrder=0"

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Accept", "application/json")
    req.Header.Set("User-Agent", "Mozilla/5.0")
    req.Header.Set("Origin", "https://www.lansfast.se")
    req.Header.Set("Referer", "https://www.lansfast.se/")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        body, _ := io.ReadAll(resp.Body)
        return nil, fmt.Errorf("non-200: %s body=%s", resp.Status, string(body))
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var data fastAPIResponseLans
    if err := json.Unmarshal(body, &data); err != nil {
        return nil, err
    }

    var houses []models.House
    for _, e := range data.Estates {
        houses = append(houses, models.House{
            Title:        e.StreetAddress,
            Price:        fmt.Sprintf("%.0f", e.StartPrice),
            SquareMeters: fmt.Sprintf("%.0f", e.Livingspace),
            Rooms:        fmt.Sprintf("%.0f", e.NumberOfRooms),
            Date:         e.PublishDate,
            Address:      e.Area,
            Image:        e.HeaderImage,
            Url:          "https://www.lansfast.se" + e.Url,
            Source:       "Länsförsäkringar",
        })
    }

    log.Println("houses lansforsakringar", len(houses))
    return houses, nil
}
