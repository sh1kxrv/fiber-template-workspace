package errors

import (
	"shared/utils/helper"

	"github.com/gofiber/fiber/v2"
)

var RepositoryError = &helper.ErrorResponse{
	StatusCode: fiber.StatusInternalServerError,
	Message:    "Unknown Repository Error",
	ServerCode: 1000,
}
