package errors

import (
	"common/utils/helper"

	"github.com/gofiber/fiber/v2"
)

var Forbidden = &helper.ErrorResponse{
	StatusCode: fiber.StatusForbidden,
	Message:    "Unknown error",
	ServerCode: 10,
}

var Unauthorized = &helper.ErrorResponse{
	StatusCode: fiber.StatusUnauthorized,
	Message:    "Unauthorized",
	ServerCode: 11,
}

var InvalidJWT = &helper.ErrorResponse{
	StatusCode: fiber.StatusUnauthorized,
	Message:    "Invalid JWT",
	ServerCode: 12,
}
