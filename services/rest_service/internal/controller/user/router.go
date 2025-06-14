package user

import (
	"rest_service/internal/config"
	"shared/middleware"

	"github.com/gofiber/fiber/v2"
)

type RouterUser struct {
	handler *UserHandler
}

func NewRouterUser(service *UserService) *RouterUser {
	return &RouterUser{
		handler: NewUserHandler(service),
	}
}

func (ru *RouterUser) RegisterRoutes(g fiber.Router) {
	userRoute := g.Group("/user", middleware.CreateJwtAuthMiddleware(config.C.Jwt.Secret))

	userRoute.Get("/me", ru.handler.GetMeHandler)
}
