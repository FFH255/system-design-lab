package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"sdl/app/internal/redis"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})

	redis.New(r, redis.Config{
		Address:  "localhost:6379",
		Username: "user",
	})

	_ = http.ListenAndServe(":3001", r)
}
