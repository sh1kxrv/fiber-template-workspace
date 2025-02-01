package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Response struct {
	Status bool `json:"status"`
	Data   any  `json:"data"`
}

type ErrorResponse struct {
	StatusCode int    `json:"code"`
	ServerCode int    `json:"serverCode"`
	Message    string `json:"message"`
}

type ServiceError struct {
	Response      *ErrorResponse
	OriginalError error
}

func NewServiceError(err error, e *ErrorResponse) *ServiceError {
	return &ServiceError{
		Response:      e,
		OriginalError: err,
	}
}

func SendError(c *fiber.Ctx, origErr error, e *ErrorResponse) error {
	resp := Response{
		Status: false,
		Data:   e,
	}
	if origErr != nil {
		logrus.Errorf("Fictive error %s, real error: %s", e.Message, origErr.Error())
	}
	return c.Status(e.StatusCode).JSON(resp)
}

func SendServiceError(c *fiber.Ctx, serr *ServiceError) error {
	return SendError(c, serr.OriginalError, serr.Response)
}

func SendSomething(c *fiber.Ctx, data any, e *ServiceError) error {
	if e != nil {
		return SendServiceError(c, e)
	}
	return SendSuccess(c, data)
}

func buildSuccessResponse(data any) Response {
	resp := Response{
		Status: true,
		Data:   data,
	}
	return resp
}

func SendSuccess(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusOK).JSON(buildSuccessResponse(data))
}
