package main

import (
	"shiro_api/database"
	"shiro_api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// CONNECTION TO DATABASE
	database.DatabaseInit()

	// MIGRATION
	// migrations.Migration()

	// FIBER INIT
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "okay",
			"age":     24,
		})
	})

	routes.RouteInit(app)

	app.Listen(":8000")
}
