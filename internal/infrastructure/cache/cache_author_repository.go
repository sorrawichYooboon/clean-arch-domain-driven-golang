package cache

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
)

type CacheAuthorRepository struct {
	Redis *redis.Client
}

func NewCacheAuthorRepository(redisClient *redis.Client) *CacheAuthorRepository {
	return &CacheAuthorRepository{Redis: redisClient}
}

func (r *CacheAuthorRepository) GetAll() ([]domain.Author, error) {
	ctx := context.Background()
	authorsJSON, err := r.Redis.Get(ctx, "authors").Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var authors []domain.Author
	if err := json.Unmarshal([]byte(authorsJSON), &authors); err != nil {
		return nil, err
	}

	return authors, nil
}

func (r *CacheAuthorRepository) SetAll(authors []domain.Author) error {
	ctx := context.Background()
	authorsJSON, err := json.Marshal(authors)
	if err != nil {
		return err
	}

	return r.Redis.Set(ctx, "authors", authorsJSON, 0).Err()
}
