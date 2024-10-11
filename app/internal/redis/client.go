package redis

import (
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Address  string `env:"REDIS_ADDRESS"`
	Username string `env:"REDIS_USERNAME"`
	Password string `env:"REDIS_PASSWORD"`
}

func NewClient(cfg Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       0,
	})
}
