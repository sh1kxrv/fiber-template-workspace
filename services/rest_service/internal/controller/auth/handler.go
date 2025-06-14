package auth

import (
	"shared/errors"
	"shared/serializer"
	"shared/utils/helper"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService *AuthService
}

func NewAuthHandler(authService *AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// @Summary Authorization
// @Tags Authorization
// @Accept json
// @Produce json
// @Param data body auth.AuthDataLogin true "Auth data"
// @Success 200 {object} auth.JwtPair
// @Failure 400 {object} helper.ErrorResponse
// @Router /api/v1/auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	data, err := serializer.GetSerializedAuthLoginData(c)

	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	pair, serr := h.authService.Login(data.Email, data.Password)
	return helper.SendSomething(c, &pair, serr)
}

// @Summary Registration
// @Tags Authorization
// @Accept json
// @Produce json
// @Param data body auth.AuthDataRegister true "Data"
// @Success 200 {object} auth.JwtPair
// @Failure 400 {object} helper.ErrorResponse
// @Router /api/v1/auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	data, err := serializer.GetSerializedAuthRegisterData(c)

	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	pair, serr := h.authService.Register(&data)
	return helper.SendSomething(c, &pair, serr)
}

// @Summary Update Access Token
// @Tags Authorization
// @Accept json
// @Produce json
// @Param data body auth.AuthDataRefresh true "Data Refresh Token"
// @Success 200 {object} auth.JwtPair
// @Failure 400 {object} helper.ErrorResponse
// @Router /api/v1/auth/refresh [post]
func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	v, err := serializer.GetSerializedAuthRefreshData(c)

	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	pair, serr := h.authService.Refresh(v.RefreshToken)
	return helper.SendSomething(c, &pair, serr)
}

func (h *AuthHandler) RegisterRoutes(g fiber.Router) {
	authRoute := g.Group("/auth")

	authRoute.Post("/login", h.Login)
	authRoute.Post("/register", h.Register)
	authRoute.Post("/refresh", h.Refresh)
}
