package home

import (
	"net/http"
)

const helloMessage = "hello"

func Handler(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(helloMessage))
}
