package main

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	elastic_search "sdl/app/internal/elastic-search"
	"sdl/app/internal/mongo"
	"sdl/app/internal/neo4j"
	"sdl/app/internal/pg"
	"sdl/app/internal/redis"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})

	// Redis
	redisClient := redis.New(r, redis.Config{
		Address:  "localhost:6379",
		Username: "user",
	})

	defer func() {
		_ = redisClient.Conn().Close()
	}()

	// mongo
	mongoClient := mongo.New(r, &mongo.Config{
		Address: "mongodb://user:user@localhost:27017",
	})

	defer func() {
		_ = mongoClient.Disconnect(context.TODO())
	}()

	// Neo4j
	neo4jClient := neo4j.New(context.TODO(), r, &neo4j.Config{
		URI:      "neo4j://localhost:7687",
		User:     "neo4j",
		Password: "user12345",
	})

	defer func() {
		_ = neo4jClient.Close(context.TODO())
	}()

	// ElasticSearch
	elastic_search.New(r, &elastic_search.Config{
		Address: "http://localhost:9200",
	})

	// PostgreSQL
	pgClient := pg.New(r, &pg.Config{
		URL: "postgres://user:user@localhost:5432/user",
	})
	defer func() {
		_ = pgClient.Close(context.TODO())
	}()

	_ = http.ListenAndServe(":3001", r)
}
