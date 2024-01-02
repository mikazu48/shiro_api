package service

import (
	"log"
	"shiro_api/database"
	"shiro_api/handler"
	"shiro_api/models"
	"shiro_api/models/converter"
	"shiro_api/request"
	"shiro_api/utils"

	"github.com/gofiber/fiber/v2"
)

type UserService struct {
}

func (c *UserService) GetList(ctx *fiber.Ctx) error {
	var users []models.User
	result := database.DB.Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(fiber.Map{
		"data": users,
	})
}
func (c *UserService) GetDataById(ctx *fiber.Ctx) error {
	var user models.User

	username := ctx.Params("username")
	if username == "" {
		handler.ResponseCustom("username invalid !", ctx)
		return nil
	}
	result := database.DB.First(&user, "username = ?", username)

	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(fiber.Map{
		"data": user,
	})
}

func (c *UserService) Create(ctx *fiber.Ctx, request *request.UserRequest) error {
	tx := database.DB.WithContext(ctx.Context()).Begin()
	defer tx.Rollback()

	if err := ctx.BodyParser(request); err != nil {
		return handler.ResponseError(ctx, err)
	}

	// Validation
	if request.FullName == "" {
		return handler.ResponseCustom("Fullname is required!", ctx)
	}

	hashedPassword, err := utils.HashingPassword(request.Password)
	if err != nil {
		return handler.ResponseError(ctx, err)
	}

	user := &models.User{
		Username: request.Username,
		Password: string(hashedPassword),
		Email:    request.Email,
		FullName: request.FullName,
	}

	err = database.DB.Create(&user).Error
	if err != nil {
		return handler.ResponseCustomErr(err.Error(), ctx, err)
	}

	return ctx.JSON(fiber.Map{
		"message": "create data successfully",
		"data":    converter.UserToResponse(user),
	})
}

func (c *UserService) Update(ctx *fiber.Ctx, request *request.UserRequest) error {
	tx := database.DB.WithContext(ctx.Context()).Begin()
	defer tx.Rollback()

	username := ctx.Params("username")
	if username == "" {
		handler.ResponseCustom("username invalid !", ctx)
		return nil
	}

	if err := ctx.BodyParser(request); err != nil {
		return handler.ResponseError(ctx, err)
	}

	// Validation
	if request.FullName == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "full name is required",
		})
	}

	hashedPassword, err := utils.HashingPassword(request.Password)
	if err != nil {
		return handler.ResponseError(ctx, err)
	}

	var user models.User
	result := database.DB.First(&user, "username = ?", username)
	if err := result.Error; err != nil {
		return handler.ResponseCustomErr("Sorry cannot found username : "+username, ctx, err)
	}

	user.Username = username
	user.Password = string(hashedPassword)
	user.Email = request.Email
	user.FullName = request.FullName

	err = database.DB.Save(&user).Error
	if err != nil {
		return handler.ResponseCustomErr(err.Error(), ctx, err)
	}

	return ctx.JSON(fiber.Map{
		"message": "update data successfully",
		"data":    converter.UserToResponse(&user),
	})
}

func (c *UserService) Delete(ctx *fiber.Ctx) error {
	tx := database.DB.WithContext(ctx.Context()).Begin()
	defer tx.Rollback()

	username := ctx.Params("username")
	if username == "" {
		handler.ResponseCustom("username invalid !", ctx)
		return nil
	}

	var user models.User
	result := database.DB.First(&user, "username = ?", username)
	if err := result.Error; err != nil {
		return handler.ResponseCustomErr("Sorry cannot found username : "+username, ctx, err)
	}

	err := database.DB.Delete(&user).Error
	if err != nil {
		return handler.ResponseCustomErr(err.Error(), ctx, err)
	}

	return ctx.JSON(fiber.Map{
		"message": "delete data successfully",
	})
}
