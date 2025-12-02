package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "house4sale/scraper"
    "log"
)

func main() {
    app := fiber.New()

    // Lägg till CORS-middleware
    app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:8081", // React Native web frontend
        AllowMethods: "GET,POST,PUT,DELETE",
    }))

    app.Get("/api/houses", func(c *fiber.Ctx) error {
        log.Println("calling /api/houses")
        data, err := scraper.Aggregate()
        if err != nil {
            return fiber.NewError(fiber.StatusInternalServerError, err.Error())
        }
        return c.JSON(data)
    })

    // Lyssna på alla interfaces
    log.Println("server listning on 8080")
    app.Listen(":8080")
}
