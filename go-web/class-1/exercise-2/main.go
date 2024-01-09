package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	// Server
	rt := chi.NewRouter()

	// Endpoints
	rt.Post("/greetings", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		var person Person
		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			log.Fatal(err)
		}

		msg := fmt.Sprintf("Hello %s %s", person.FirstName, person.LastName)

		_, err = w.Write([]byte(msg))

		if err != nil {
			log.Fatal(err)
		}
	})

	fmt.Println("Server is up!")
	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}
