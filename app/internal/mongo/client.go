package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"time"
)

type Config struct {
	Address string
}

func NewClient(cfg *Config) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(cfg.Address).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, err
	}

	return client, nil
}
