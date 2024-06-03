package main

import (
	"go-auth-token/helper"
	"go-auth-token/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	// database.OpenConnection()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8000",
		AllowCredentials: true,
	}))

	routes.Setup(app)

	err := app.Listen("localhost:8000")
	helper.PanicIfError(err)
}
