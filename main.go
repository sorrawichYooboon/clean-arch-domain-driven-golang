package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/config"
	_ "github.com/sorrawichYooboon/clean-arch-domain-driven-golang/docs"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/infrastructure/cache"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/infrastructure/database"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/infrastructure/http"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/infrastructure/http/controller"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/usecase"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Book Store Management API
// @version 1.0
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
// @type Bearer
// @description You need to register and login on that api and Provide the JWT token prefixed with "Bearer " (including a space). For example: "Bearer your_token_here"

// @title Book Store Management API
// @version 1.0
// @securityDefinitions.bearer BearerAuth
// @in header
// @name Authorization
// @type string
// @description This API enables management of a book store, allowing users to perform operations such as adding, updating, retrieving, and deleting books.
// You must register and log in to receive a JWT token. Provide this token in the Authorization header, prefixed with "Bearer " (including a space).
// Example: "Bearer your_token_here".
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	e := echo.New()

	cfg := config.Initialize()

	bookRepo := database.NewBookRepository(cfg.DB)
	cacheBookRepo := cache.NewCacheBookRepository(cfg.Redis)
	bookUseCase := usecase.NewBookUseCase(bookRepo, *cacheBookRepo)
	bookHandler := controller.NewBookHandler(bookUseCase)

	authorRepo := database.NewAuthorRepository(cfg.DB)
	cacheAuthorRepo := cache.NewCacheAuthorRepository(cfg.Redis)
	authorUseCase := usecase.NewAuthorUseCase(authorRepo, *cacheAuthorRepo)
	authorHandler := controller.NewAuthorHandler(authorUseCase)

	userRepo := database.NewUserRepository(cfg.DB)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := controller.NewUserHandler(userUseCase, cfg.SecretKey)

	http.SetupUserRoutes(e, userHandler)
	http.SetupBookRoutes(e, bookHandler, cfg.SecretKey)
	http.SetupAuthorRoutes(e, authorHandler, cfg.SecretKey)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
