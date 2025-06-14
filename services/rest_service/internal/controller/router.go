package controller

import (
	"rest_service/internal/controller/auth"
	"rest_service/internal/controller/user"
	"shared/driver/mongodb"
	"shared/driver/mongodb/repository"

	"github.com/gofiber/fiber/v2"
)

func makeBaseGroup(app *fiber.App) fiber.Router {
	api := app.Group("/api")

	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})

	return v1
}

func InitRouter(app *fiber.App, db *mongodb.MongoInstance) {
	v1 := makeBaseGroup(app)

	// Repositories
	userRepository := repository.NewUserRepository(db)

	// Services
	userService := user.NewUserService(userRepository)
	authService := auth.NewAuthService(userRepository)

	// Routers
	authRouter := auth.NewRouterAuth(authService)
	userRouter := user.NewRouterUser(userService)

	// Registration
	authRouter.RegisterRoutes(v1)
	userRouter.RegisterRoutes(v1)
}
