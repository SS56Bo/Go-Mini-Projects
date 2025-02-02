package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Sup, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
