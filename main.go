package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hojin-kr/account"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/account/:token", func(c *fiber.Ctx) error {
		account := account.Account{}
		account.Token = c.Params("token")
		account.GetAccount()
		return c.JSON(account)
	})

	app.Get("/account/new/:token", func(c *fiber.Ctx) error {
		account := account.Account{}
		account.Token = c.Params("token")
		// get account if not exist new account
		account.GetAccount()
		if account.UUID != "" {
			return c.JSON(account)
		}
		account.NewAccount()
		return c.JSON(account)
	})

	app.Listen(":3000")
}
