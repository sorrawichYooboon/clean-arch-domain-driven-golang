package repository

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
)

type CacheBookRepository struct {
	Redis *redis.Client
}

func NewCacheBookRepository(redisClient *redis.Client) *CacheBookRepository {
	return &CacheBookRepository{Redis: redisClient}
}

func (r *CacheBookRepository) GetAll() ([]domain.Book, error) {
	ctx := context.Background()
	booksJSON, err := r.Redis.Get(ctx, "books").Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var books []domain.Book
	if err := json.Unmarshal([]byte(booksJSON), &books); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *CacheBookRepository) SetAll(books []domain.Book) error {
	ctx := context.Background()
	booksJSON, err := json.Marshal(books)
	if err != nil {
		return err
	}

	return r.Redis.Set(ctx, "books", booksJSON, 0).Err()
}
