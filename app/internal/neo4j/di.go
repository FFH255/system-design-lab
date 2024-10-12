package neo4j

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func New(ctx context.Context, router *mux.Router, cfg *Config) neo4j.DriverWithContext {
	driver := NewClient(ctx, cfg)
	ServeHTTP(router, NewHandler(NewRepository(driver)))
	return driver
}
