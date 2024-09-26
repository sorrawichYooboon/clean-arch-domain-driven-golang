package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/sorrawichYooboon/clean-arch-domain-driven-golang/docs"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/infrastructure/cache"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/infrastructure/config"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/infrastructure/database"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/infrastructure/http"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/infrastructure/http/controller.go"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/usecase"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	e := echo.New()

	db, redisClient := config.Initialize()
	bookRepo := database.NewBookRepository(db)
	cacheBookRepo := cache.NewCacheBookRepository(redisClient)
	bookUseCase := usecase.NewBookUseCase(bookRepo, *cacheBookRepo)
	bookHandler := controller.NewBookHandler(bookUseCase)

	authorRepo := database.NewAuthorRepository(db)
	cacheAuthorRepo := cache.NewCacheAuthorRepository(redisClient)
	authorUseCase := usecase.NewAuthorUseCase(authorRepo, *cacheAuthorRepo)
	authorHandler := controller.NewAuthorHandler(authorUseCase)

	http.SetupBookRoutes(e, bookHandler)
	http.SetupAuthorRoutes(e, authorHandler)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
