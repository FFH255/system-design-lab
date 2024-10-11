package redis

import (
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

func New(router *mux.Router, cfg Config) *redis.Client {
	client := NewClient(cfg)
	ServeHTTP(router, NewHandler(NewService(NewRepository(client))))
	return client
}
