package models

type House struct {
    Title       string `json:"title"`
    Price       string `json:"price"`
    Address     string `json:"address"`
    Image       string `json:"image"`
    Url         string `json:"url"`
    Source      string `json:"source"`
}
