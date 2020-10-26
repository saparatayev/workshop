package main

import (
	"log"
	"net/http"
	"workshop/internal/handler"

	"github.com/go-chi/chi"
)

func main() {

	h := handler.NewHandler()

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	log.Print("starting server")

	log.Fatal(http.ListenAndServe(":8081", r))

	log.Print("shutting server down")
}
