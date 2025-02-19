package errors

import (
	"shared/utils/helper"

	"github.com/gofiber/fiber/v2"
)

var JwtPairGenerationError = &helper.ErrorResponse{
	StatusCode: fiber.StatusInternalServerError,
	Message:    "JWT Pair Generation Error",
	ServerCode: 1100,
}

var JwtRefreshTokenInvalid = &helper.ErrorResponse{
	StatusCode: fiber.StatusUnauthorized,
	Message:    "Invalid Refresh Token",
	ServerCode: 1101,
}
