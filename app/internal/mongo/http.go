package mongo

import (
	"github.com/gorilla/mux"
)

func ServeHTTP(defaultRouter *mux.Router, handler *Handler) {
	r := defaultRouter.PathPrefix("/mongo/group").Subrouter()

	r.HandleFunc("", handler.Create).Methods("POST")
	r.HandleFunc("/{id}", handler.Get).Methods("GET")
	r.HandleFunc("/{id}", handler.Update).Methods("PUT")
	r.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
