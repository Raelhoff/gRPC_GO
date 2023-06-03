package main

import (
    "github.com/gofiber/fiber/v2"
    "fmt"
    "log"
)

type Data struct {
    version uint8 `json:"version"`
    id      uint8 `json:"id"` 
 }

func DataHandler(c *fiber.Ctx) error {
    var p Data
    if err := c.BodyParser(&amp;p); err != nil {
        return err
    }

    return c.SendString(fmt.Sprintf("version: %d - id: %d\n", p.version, p.id))
}

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
    })

   app.Get("/lora",  func(c *fiber.Ctx) error {
        fmt.Printf("Get - Chegou mensagem")
	return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
    })

    
    app.Post("/lora", DataHandler)

    app.Listen(":3033")
}
