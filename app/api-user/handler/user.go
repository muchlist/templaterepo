package handler

import (
	"monorepo/business/user"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	usecase *user.UserService
}

func NewUserHandler(usecase *user.UserService) UserHandler {
	return UserHandler{
		usecase: usecase,
	}
}

func (u *UserHandler) Login(c *fiber.Ctx) error {

	username := c.Get("username")
	password := c.Get("password")

	user, _ := u.usecase.Login(c.Context(), username, password)
	return c.JSON(fiber.Map{
		"data":  user,
		"error": nil,
	})
}
