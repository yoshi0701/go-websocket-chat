package main

import (
	"log"
	"net/http"
	handlers "ws/cmd/internal"
)

func main() {
	mux := routes()

	log.Println("starting channel listener")
	go handlers.ListenToWsChannel()

	log.Println("Starting web server on port 8080")

	_ = http.ListenAndServe(":8080", mux)
}
