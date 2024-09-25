package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/config"
	_ "github.com/sorrawichYooboon/clean-arch-domain-driven-golang/docs"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/delivery/http"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/delivery/routes"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/repository"
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

	bookRepo := repository.NewBookRepository(db)
	cacheBookRepo := repository.NewCacheBookRepository(redisClient)
	bookUseCase := usecase.NewBookUseCase(bookRepo, cacheBookRepo)
	bookHandler := http.NewBookHandler(bookUseCase)

	authorRepo := repository.NewAuthorRepository(db)
	cacheAuthorRepo := repository.NewCacheAuthorRepository(redisClient)
	authorUseCase := usecase.NewAuthorUseCase(authorRepo, cacheAuthorRepo)
	authorHandler := http.NewAuthorHandler(authorUseCase)

	routes.SetupBookRoutes(e, bookHandler)
	routes.SetupAuthorRoutes(e, authorHandler)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
