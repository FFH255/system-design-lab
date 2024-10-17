package pg

import "github.com/gorilla/mux"

func ServeHTTP(defaultRouter *mux.Router, handler *Handler) {
	r := defaultRouter.PathPrefix("/pg/attendance").Subrouter()

	r.HandleFunc("", handler.Create).Methods("POST")
	r.HandleFunc("", handler.GetAll).Methods("GET")
	r.HandleFunc("/{id}", handler.Update).Methods("PUT")
	r.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
