package utils

import (
	"bytes"
	"common/validator"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetSerializedBodyData[T any](c *fiber.Ctx) (T, error) {
	var data T

	rawBody := c.Body()

	decoder := json.NewDecoder(bytes.NewReader(rawBody))
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&data); err != nil {
		return data, err
	}

	validator := validator.GetValidatorInstance()
	if err := validator.Struct(data); err != nil {
		return data, err
	}

	return data, nil
}

func GetJwtUserLocalWithParsedID(c *fiber.Ctx) (JwtClaims, primitive.ObjectID, error) {
	localUser := c.Locals("user")
	if localUser == nil {
		return JwtClaims{}, primitive.NilObjectID, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	parsedLocalUser, ok := localUser.(JwtClaims)
	if !ok {
		return JwtClaims{}, primitive.NilObjectID, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	parsedID, err := primitive.ObjectIDFromHex(parsedLocalUser.ID)
	if err != nil {
		return JwtClaims{}, primitive.NilObjectID, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	return parsedLocalUser, parsedID, nil
}
