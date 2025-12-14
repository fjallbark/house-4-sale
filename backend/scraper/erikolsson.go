package scraper

import (
	"context"
	"house4sale/models"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func ScrapeErikOlsson() ([]models.House, error) {
    log.Println("ScrapeErikOlsson")
opts := append(chromedp.DefaultExecAllocatorOptions[:],
    chromedp.ExecPath("chromium"),
    chromedp.Headless,
    chromedp.DisableGPU,
    chromedp.NoSandbox,
)



    allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
    defer cancel()

    ctx, cancel := chromedp.NewContext(allocCtx)
    defer cancel()

    ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
    defer cancel()

    var htmlContent string
    url := "https://www.erikolsson.se/homes?rooms_min=5&label=municipality%3A110%3AH%C3%B6%C3%B6r&municipalityIds=110"

    err := chromedp.Run(ctx,
        chromedp.Navigate(url),
        chromedp.Sleep(3*time.Second), // vänta tills sidan laddat JS
        chromedp.OuterHTML("html", &htmlContent),
    )
    if err != nil {
        return nil, err
    }

    var houses []models.House

    // Enkel parsing med strings (kan ersättas med goquery om du vill)
    entries := strings.Split(htmlContent, `search-hit__link`)
    for _, e := range entries[1:] { // hoppa över första delen
        var h models.House

        // Title / Address
        start := strings.Index(e, `">`)
        end := strings.Index(e, `</a>`)
        if start > -1 && end > start {
            h.Title = strings.TrimSpace(e[start+2 : end])
        }

        // Price
        priceStart := strings.Index(e, `search-hit__info--spacer-se sweden--price notranslate">`)
        if priceStart > -1 {
            priceEnd := strings.Index(e[priceStart:], `</span>`)
            if priceEnd > -1 {
                h.Price = strings.TrimSpace(e[priceStart+50 : priceStart+priceEnd])
            }
        }

        // URL
        urlStart := strings.Index(e, `href="`)
        if urlStart > -1 {
            urlEnd := strings.Index(e[urlStart+6:], `"`)
            if urlEnd > -1 {
                h.Url = "https://www.erikolsson.se" + e[urlStart+6:urlStart+6+urlEnd]
            }
        }

        h.Source = "ErikOlsson"
        houses = append(houses, h)
    }

    log.Println("houses ErikOlsson:", len(houses))
    return houses, nil
}
