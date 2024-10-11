package api

import (
	"encoding/json"
	"net/http"
)

func ReadJSON[T any](r *http.Request) (T, error) {
	var body T
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	return body, err
}
