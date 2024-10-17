package pg

import (
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

func New(router *mux.Router, cfg *Config) *pgx.Conn {
	client := NewClient(cfg)

	ServeHTTP(router, NewHandler(NewRepository(client)))

	return client
}
