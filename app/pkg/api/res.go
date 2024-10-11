package api

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

func NewError(message string) Error {
	return Error{message}
}

func WriteJSON(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(res)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, NewError(message))
}
