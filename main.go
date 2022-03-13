package main

import (
	"go-microservice-in-30-min/home"
	"go-microservice-in-30-min/server"
	"log"
	"net/http"
	"os"
)

var (
	AppServiceAddress = getenv("APP_SERVICE_ADDRESS", "localhost:8080")
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home.Handler)
	srv := server.New(mux, AppServiceAddress)
	//err := http.ListenAndServe(":8080", mux) has high timeouts=10m
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start the application", err)
	}
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
