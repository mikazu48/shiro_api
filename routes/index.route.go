package routes

import (
	"shiro_api/controllers"

	"github.com/gofiber/fiber/v2"
)

// DECLARE CONTROLLER
var authController = controllers.NewAuthController()
var userController = controllers.NewUserController()

func RouteInit(app *fiber.App) {
	app.Post("/login", authController.Login)
	// app.Get("/me", authController.AuthToken, authController.GetMe)
	app.Get("/users", authController.AuthToken, userController.GetList)
	app.Get("/users/:username", authController.AuthToken, userController.GetList)
	app.Post("/users", authController.AuthToken, userController.Create)
	app.Put("/users/:username", authController.AuthToken, userController.Update)
	app.Delete("/users/:username", authController.AuthToken, userController.Delete)
}
