package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// APIError is an internal error to json
type APIError struct {
	Code    int    `json:"code"`
	API     string `json:"api"`
	Message string `json:"message"`
}

// NewAPIError returns a struct with error
func NewAPIError(code int, api, msg string) *APIError {
	return &APIError{
		Code:    code,
		API:     api,
		Message: msg,
	}
}

// JSONRender converts an object to its json string and prints it
func JSONRender(obj interface{}, w http.ResponseWriter) {
	js, err := json.Marshal(obj)
	if err != nil {
		errorRender(500, "json_render", err.Error(), w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", js)
}

func errorRender(code int, api, msg string, w http.ResponseWriter) {
	e := NewAPIError(code, api, msg)
	es, err := json.Marshal(e)
	if err != nil {
		errorRender(code, "errorRender", err.Error(), w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", es)
}
