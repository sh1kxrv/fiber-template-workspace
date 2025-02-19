package param

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ParamPrimitiveID(c *fiber.Ctx, key string) (primitive.ObjectID, error) {
	id := c.Params(key)
	parsedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return parsedId, nil
}
