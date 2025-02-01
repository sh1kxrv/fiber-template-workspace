package param

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetLimitOffset(c *fiber.Ctx) (limit int64, offset int64) {
	limitRaw := c.Query("limit", "10")
	offsetRaw := c.Query("offset", "0")

	limitInt, err := strconv.ParseInt(limitRaw, 10, 64)
	if err != nil {
		limitInt = 10
	}

	offsetInt, err := strconv.ParseInt(offsetRaw, 10, 64)
	if err != nil {
		offsetInt = 0
	}

	return limitInt, offsetInt
}
