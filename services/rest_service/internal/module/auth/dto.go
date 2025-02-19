package auth

import (
	"shared/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthDataLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthDataRegister struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstName" validate:"required,min=2,max=50"`
	LastName  string `json:"lastName" validate:"required,min=2,max=50"`
}

type AuthDataRefresh struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

func GetSerializedAuthRefreshData(c *fiber.Ctx) (AuthDataRefresh, error) {
	return utils.GetSerializedBodyData[AuthDataRefresh](c)
}

func GetSerializedAuthLoginData(c *fiber.Ctx) (AuthDataLogin, error) {
	return utils.GetSerializedBodyData[AuthDataLogin](c)
}

func GetSerializedAuthRegisterData(c *fiber.Ctx) (AuthDataRegister, error) {
	return utils.GetSerializedBodyData[AuthDataRegister](c)
}
