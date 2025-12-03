package models

type House struct {
    Title           string `json:"title"`
    SquareMeters    string `json:"squareMeters"`
    Rooms           string `json:"rooms"`
    Date            string `json:"Date"`
    Price           string `json:"price"`
    Address         string `json:"address"`
    Image           string `json:"image"`
    Url             string `json:"url"`
    Source          string `json:"source"`
}
