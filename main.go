package main

import (
	"github.com/LitHaxor/fiber-go-gorm.git/database"
	"github.com/gofiber/fiber/v2"
)

type Data struct {
    name string `json:"name"`
    age uint8 `json:"age"`
}


func main() {
    // Start a new fiber app
    app := fiber.New()

   database.ConnectDB()
    // Send a string back for GET calls to the endpoint "/"
    app.Get("/", func(c *fiber.Ctx) error {
        data := []Data{{
            name: "hello",
            age: 20,
        }}
        return c.JSON(&data)
    })

    // Listen on PORT 3000
    app.Listen(":3000")
}