package elastic_search

import (
	"github.com/gorilla/mux"
)

func New(router *mux.Router, cfg *Config) {
	client := NewClient(cfg)
	ServeHTTP(router, NewHandler(NewRepository(client)))
}
