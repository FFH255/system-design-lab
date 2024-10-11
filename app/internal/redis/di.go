package redis

import "github.com/gorilla/mux"

func New(router *mux.Router, cfg Config) {
	ServeHTTP(router, NewHandler(NewService(NewRepository(NewClient(cfg)))))
}
