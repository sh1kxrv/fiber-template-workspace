package errors

import (
	"shared/utils/helper"

	"github.com/gofiber/fiber/v2"
)

var EntityAlreadyExists = &helper.ErrorResponse{
	StatusCode: fiber.StatusConflict,
	Message:    "Entity Already Exists",
	ServerCode: 1200,
}

var ParseIDError = &helper.ErrorResponse{
	StatusCode: fiber.StatusInternalServerError,
	Message:    "Parse ID Error",
	ServerCode: 1201,
}

var EntityNotExists = &helper.ErrorResponse{
	StatusCode: fiber.StatusNotFound,
	Message:    "Entity Not Exists",
	ServerCode: 1202,
}
