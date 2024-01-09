package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	// Server
	rt := chi.NewRouter()

	// Endpoints
	rt.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("pong"))

		if err != nil {
			panic(err)
		}
	})

	fmt.Println("Server is starting...")
	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}
