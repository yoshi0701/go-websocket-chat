package main

import (
	"github.com/bmizerany/pat"
	"net/http"
	handlers "ws/cmd/internal"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))

	return mux
}
