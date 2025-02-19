package handler

import "github.com/labstack/echo/v4"

type BookHandler interface {
	GetAll(echo.Context) error
	Create(echo.Context) error
	Update(echo.Context) error
	Delete(c echo.Context) error
}

type AuthorHandler interface {
	GetAll(echo.Context) error
	Create(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}

type UserHandler interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}
