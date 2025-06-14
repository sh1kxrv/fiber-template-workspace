package serializer

import (
	"shared/transfer/dto"
	"shared/utils"

	"github.com/gofiber/fiber/v2"
)

func GetSerializedAuthRefreshData(c *fiber.Ctx) (dto.AuthDataRefresh, error) {
	return utils.GetSerializedBodyData[dto.AuthDataRefresh](c)
}

func GetSerializedAuthLoginData(c *fiber.Ctx) (dto.AuthDataLogin, error) {
	return utils.GetSerializedBodyData[dto.AuthDataLogin](c)
}

func GetSerializedAuthRegisterData(c *fiber.Ctx) (dto.AuthDataRegister, error) {
	return utils.GetSerializedBodyData[dto.AuthDataRegister](c)
}
