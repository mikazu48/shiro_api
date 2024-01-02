package handler

import (
	"github.com/gofiber/fiber/v2"
)

func ResponseError(ctx *fiber.Ctx, err interface{}) error {
	return ctx.Status(400).JSON(fiber.Map{
		"message": "There an error, please check again!",
	})
}

func ResponseNotFound(model string, ctx *fiber.Ctx, err interface{}) error {
	return ctx.Status(404).JSON(fiber.Map{
		"message": model + " Not Found",
	})
}

func ResponseCustomErr(message string, ctx *fiber.Ctx, err interface{}) error {
	return ctx.Status(405).JSON(fiber.Map{
		"message": message,
	})
}

func ResponseCustom(message string, ctx *fiber.Ctx) error {
	return ctx.Status(405).JSON(fiber.Map{
		"message": message,
	})
}
