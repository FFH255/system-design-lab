package main

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"sdl/app/internal/mongo"
	"sdl/app/internal/neo4j"
	"sdl/app/internal/redis"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})

	redisClient := redis.New(r, redis.Config{
		Address:  "localhost:6379",
		Username: "user",
	})

	defer func() {
		_ = redisClient.Conn().Close()
	}()

	mongoClient := mongo.New(r, &mongo.Config{
		Address: "mongodb://user:user@localhost:27017",
	})

	defer func() {
		_ = mongoClient.Disconnect(context.TODO())
	}()

	neo4jClient := neo4j.New(context.TODO(), r, &neo4j.Config{
		URI:      "neo4j://localhost:7687",
		User:     "neo4j",
		Password: "user12345",
	})

	defer func() {
		_ = neo4jClient.Close(context.TODO())
	}()

	_ = http.ListenAndServe(":3001", r)
}
