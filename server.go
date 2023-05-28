package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	setupRoutes(app)

	err := app.Listen(":8080")

	if err != nil {
		panic(err)
	}
}
