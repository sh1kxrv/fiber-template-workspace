package auth

import "github.com/gofiber/fiber/v2"

type RouterAuth struct {
	handler *AuthHandler
}

func NewRouterAuth(service *AuthService) *RouterAuth {
	return &RouterAuth{
		handler: NewAuthHandler(service),
	}
}

func (ra *RouterAuth) RegisterRoutes(g fiber.Router) {
	authRoute := g.Group("/auth")

	authRoute.Post("/login", ra.handler.Login)
	authRoute.Post("/register", ra.handler.Register)
	authRoute.Post("/refresh", ra.handler.Refresh)
}
