package errors

import (
	"common/utils/helper"

	"github.com/gofiber/fiber/v2"
)

var UnknownError = &helper.ErrorResponse{
	StatusCode: fiber.StatusInternalServerError,
	Message:    "Unknown error",
	ServerCode: 0,
}
