package main

import (
	"go_auth_token/app"
	"go_auth_token/helper"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app.NewDB()

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	err := app.Listen("localhost:8000")
	helper.PanicIfError(err)
}
