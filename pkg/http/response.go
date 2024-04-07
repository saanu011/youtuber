package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	headerContentType = "Content-Type"
	contentTypeJSON   = "application/json"
)

func JSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set(headerContentType, contentTypeJSON)
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(body); err != nil {
		fmt.Println("unable to write to body", err)
	}

	if isSuccess(statusCode) {
		fmt.Println("http.request.success")
	} else {
		fmt.Println("http.request.error")
	}
}
