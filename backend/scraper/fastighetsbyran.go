package scraper

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"house4sale/models"
)

type fastAPIResponse struct {
    Results []fastItem `json:"results"`
}

type fastItem struct {
    StorRubrik string   `json:"storRubrik"`
    LitenRubrik string  `json:"litenRubrik"`
    MetaData    []string `json:"metaData"`
    BildUrl     string   `json:"bildUrl"`
    Url         string   `json:"url"`
}

func ScrapeFastighetsbyran() ([]models.House, error) {
    log.Println("ScrapeFastighetsbyran")

    // ---- Payload ----
    payload := map[string]any{
        "inkluderaNyproduktion":    false,
        "inkluderaPaaGaang":        true,
        "positioner":               []int{},
        "valdaKommuner":            []int{},
        "valdaKontor":              []int{},
        "valdaLaen":                []int{},
        "valdaMaeklarObjektTyper":  []int{0},
        "valdaNaeromraaden":        []int{17501},
        "valdaNyckelord":           []string{},
        "valdaPostorter":           []string{},
        "spraak":                   "sv",
    }

    jsonPayload, err := json.Marshal(payload)
    if err != nil {
        return nil, err
    }

    // ---- Request ----
    req, err := http.NewRequest("POST",
        "https://www.fastighetsbyran.com/HemsidanAPI/api/v1/soek/objekt/1/false/",
        bytes.NewReader(jsonPayload),
    )
    if err != nil {
        return nil, err
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept", "application/json")
    req.Header.Set("User-Agent", "Mozilla/5.0")
    req.Header.Set("Accept-Language", "sv-SE,sv;q=0.9")
    req.Header.Set("Origin", "https://www.fastighetsbyran.com")
    req.Header.Set("Referer", "https://www.fastighetsbyran.com/sv/sverige/till-salu")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, errors.New("non-200 response: " + resp.Status)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    log.Println("RAW RESPONSE:", string(body))

    // ---- Parse API response ----
    var data fastAPIResponse
    if err := json.Unmarshal(body, &data); err != nil {
        log.Println("Unmarshal error:", err)
        return nil, err
    }

    // ---- Convert to []models.House ----
    var houses []models.House

for _, item := range data.Results {
    price := ""
    if len(item.MetaData) > 0 {
        price = item.MetaData[0] // t.ex. "4 695 000 kr"
    }

    houses = append(houses, models.House{
        Title:   item.StorRubrik,
        Price:   price,
        Address: item.LitenRubrik,
        Image:   item.BildUrl,
        Url:     item.Url, // redan full URL i API:et
        Source:  "Fastighetsbyrån",
    })
}


    log.Println("houses", houses)

    return houses, nil
}
