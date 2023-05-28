package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/svetamazz/currency-exchange-rate/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/rate", handlers.GetCurrencyRate)
	
	app.Post("/sendEmails", handlers.SendEmailsToSubscribers)

	app.Post("/subscribe", handlers.SubscribeEmail)
}
