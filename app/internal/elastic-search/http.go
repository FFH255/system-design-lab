package elastic_search

import "github.com/gorilla/mux"

func ServeHTTP(defaultRouter *mux.Router, handler *Handler) {
	r := defaultRouter.PathPrefix("/elastic/course").Subrouter()

	r.HandleFunc("", handler.Create).Methods("POST")
	r.HandleFunc("/search", handler.Search).Methods("GET")
	r.HandleFunc("/{id}", handler.Update).Methods("PUT")
	r.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
