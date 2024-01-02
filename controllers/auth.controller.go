package controllers

import (
	"shiro_api/database"
	"shiro_api/handler"
	"shiro_api/models"
	"shiro_api/request"
	"shiro_api/utils"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	request := new(request.LoginRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		return handler.ResponseError(ctx, err)
	}

	validate := validator.New()
	errorValidate := validate.Struct(request)
	if errorValidate != nil {
		return handler.ResponseCustomErr("Invalid validation struct", ctx, err)
	}

	var user models.User
	err = database.DB.First(&user, "username =?", request.Username).Error

	if err != nil {
		return handler.ResponseNotFound("User", ctx, err)
	}

	isValid := utils.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		return handler.ResponseCustomErr("Password is invalid", ctx, err)
	}

	claims := jwt.MapClaims{}
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["fullname"] = user.FullName
	claims["exp"] = time.Now().AddDate(0, 0, 14).Unix()

	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil {
		return handler.ResponseCustomErr("Generate token is invalid", ctx, err)
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}

func (c *AuthController) AuthToken(ctx *fiber.Ctx) error {
	var tokenString string
	authorization := ctx.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if ctx.Cookies("token") != "" {
		tokenString = ctx.Cookies("token")
	}

	if tokenString == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	_, err := utils.VerifyToken(tokenString)
	if err != nil {
		return handler.ResponseCustomErr("Token is invalid", ctx, err)
	}

	return ctx.Next()

}
