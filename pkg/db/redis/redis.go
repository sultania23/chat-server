package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type Config struct {
	Host     string
	Port     string
	DB       int
	Password string
	Expires  time.Duration
}

func NewRedisClient(cfg Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})
}
