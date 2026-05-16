package ui

import (
	"encoding/json"
	"net/http"
)

// ParseJSON reads JSON from the request body and unmarshals it into the provided type T.
// It returns an error if the body is invalid JSON or cannot be decoded to type T.
func ParseJSON[T any](r *http.Request) (T, error) {
	var data T
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	return data, err
}
