package mongo

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func New(router *mux.Router, cfg *Config) *mongo.Client {
	client, err := NewClient(cfg)

	if err != nil {
		panic(err)
	}

	ServeHTTP(router, NewHandler(NewService(NewRepository(client))))

	return client
}
