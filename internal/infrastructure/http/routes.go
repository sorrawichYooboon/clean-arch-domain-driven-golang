package http

import (
	"github.com/labstack/echo/v4"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/infrastructure/http/handler"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/infrastructure/middleware"
)

func SetupBookRoutes(e *echo.Echo, bookHandler handler.BookHandler, secretKey string) {
	bookGroup := e.Group("/books")
	bookGroup.Use(middleware.AuthMiddleware(secretKey))

	bookGroup.GET("", bookHandler.GetAll)
	bookGroup.POST("", bookHandler.Create)
	bookGroup.PUT("/:id", bookHandler.Update)
	bookGroup.DELETE("/:id", bookHandler.Delete)
}

func SetupAuthorRoutes(e *echo.Echo, authorHandler handler.AuthorHandler, secretKey string) {
	authorGroup := e.Group("/authors")
	authorGroup.Use(middleware.AuthMiddleware(secretKey))

	authorGroup.GET("", authorHandler.GetAll)
	authorGroup.POST("", authorHandler.Create)
	authorGroup.PUT("/:id", authorHandler.Update)
	authorGroup.DELETE("/:id", authorHandler.Delete)
}

func SetupUserRoutes(e *echo.Echo, userHandler handler.UserHandler) {
	e.POST("/register", userHandler.Register)
	e.POST("/login", userHandler.Login)
}
