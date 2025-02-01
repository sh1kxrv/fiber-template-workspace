package api

import "github.com/gofiber/fiber/v2"

type Router interface {
	RegisterRoutes(g fiber.Router)
}
