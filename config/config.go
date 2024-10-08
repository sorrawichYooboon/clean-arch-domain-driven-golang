package config

import (
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DB        *gorm.DB
	Redis     *redis.Client
	SecretKey string
}

func Initialize() *Config {
	db, err := initializeDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	redisClient, err := initializeRedis()
	if err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}

	secretKey := getEnv("SECRET_KEY", "default_secret_key")

	return &Config{
		DB:        db,
		Redis:     redisClient,
		SecretKey: secretKey,
	}
}

func initializeDatabase() (*gorm.DB, error) {
	dbHost := getEnv("POSTGRES_HOST", "localhost")
	dbUser := getEnv("POSTGRES_USER", "")
	dbPassword := getEnv("POSTGRES_PASSWORD", "")
	dbName := getEnv("POSTGRES_DB", "")
	dbPort := getEnv("POSTGRES_PORT", "5432")

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initializeRedis() (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: getEnv("REDIS_ADDR", "localhost:6379"),
	})

	ctx := context.Background()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return redisClient, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
