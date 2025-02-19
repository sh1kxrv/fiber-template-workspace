package errors

import (
	"shared/utils/helper"

	"github.com/gofiber/fiber/v2"
)

var UnknownError = &helper.ErrorResponse{
	StatusCode: fiber.StatusInternalServerError,
	Message:    "Unknown error",
	ServerCode: 0,
}

var BadRequest = &helper.ErrorResponse{
	StatusCode: fiber.StatusBadRequest,
	Message:    "Bad request",
	ServerCode: 1,
}

var CryptoError = &helper.ErrorResponse{
	StatusCode: fiber.StatusInternalServerError,
	Message:    "Crypto Error",
	ServerCode: 2,
}
