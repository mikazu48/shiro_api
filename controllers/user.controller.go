package controllers

import (
	"shiro_api/request"
	"shiro_api/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) GetList(ctx *fiber.Ctx) error {
	return c.UserService.GetList(ctx)
}

func (c *UserController) GetDataById(ctx *fiber.Ctx) error {
	return c.UserService.GetDataById(ctx)
}

func (c *UserController) Create(ctx *fiber.Ctx) error {
	request := new(request.UserRequest)
	return c.UserService.Create(ctx, request)
}

func (c *UserController) Update(ctx *fiber.Ctx) error {
	request := new(request.UserRequest)
	return c.UserService.Update(ctx, request)
}

func (c *UserController) Delete(ctx *fiber.Ctx) error {
	return c.UserService.Delete(ctx)
}
