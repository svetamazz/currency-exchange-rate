package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/svetamazz/currency-exchange-rate/helpers"
)

func GetCurrencyRate(c *fiber.Ctx) error {
	rate, err := helpers.GetExchangeRate()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get exchange rate")
	}

	return c.JSON(rate.BTCtoUAH)
}

func SubscribeEmail(c *fiber.Ctx) error {
	email := c.FormValue("email")
	if email == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Email is required")
	}

	emails, err := helpers.GetEmails()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get emails")
	}

	if helpers.Contains(emails, email) {
		return c.Status(fiber.StatusConflict).SendString("Email already exists")
	}

	emails.Emails = append(emails.Emails, email)

	err = helpers.SaveEmails(emails)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to save email")
	}

	return c.SendString("Email saved successfully")
}

func SendEmailsToSubscribers(c *fiber.Ctx) error {
	rate, err := helpers.GetExchangeRate()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get exchange rate")
	}

	emails, err := helpers.GetEmails()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get emails")
	}

	if err := helpers.SendEmails(emails.Emails, rate.BTCtoUAH); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to send emails")
	}

	return c.SendString("Emails sent successfully")
}