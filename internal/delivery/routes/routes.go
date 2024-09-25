package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/delivery/http"
)

func SetupBookRoutes(e *echo.Echo, bookHandler *http.BookHandler) {
	e.GET("/books", bookHandler.GetAll)
	e.POST("/books", bookHandler.Create)
	e.PUT("/books/:id", bookHandler.Update)
	e.DELETE("/books/:id", bookHandler.Delete)
}

func SetupAuthorRoutes(e *echo.Echo, authorHandler *http.AuthorHandler) {
	e.GET("/authors", authorHandler.GetAll)
	e.POST("/authors", authorHandler.Create)
	e.PUT("/authors/:id", authorHandler.Update)
	e.DELETE("/authors/:id", authorHandler.Delete)
}
