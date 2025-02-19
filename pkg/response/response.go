package response

import (
	"github.com/labstack/echo/v4"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/pkg/apperror"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/pkg/httperror"
)

type APIResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c echo.Context, statusCode int, response Response, data interface{}) error {
	return c.JSON(statusCode, APIResponse{
		Code:    response.Code,
		Message: response.Message,
		Data:    data,
	})
}

func Error(c echo.Context, statusCode int, err error) error {
	var httpError httperror.HTTPError

	switch e := err.(type) {
	case *apperror.UseCaseError:
		httpError = httperror.HTTPError{
			Code:    e.Code,
			Message: e.Message,
		}
	case *httperror.HTTPError:
		httpError = *e
	default:
		httpError = httperror.ErrInternalServer
	}

	return c.JSON(statusCode, APIResponse{
		Code:    httpError.Code,
		Message: httpError.Message,
	})
}
