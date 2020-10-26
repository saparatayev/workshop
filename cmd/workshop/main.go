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

	log.Fatal(http.ListenAndServe(":8081", r))
}
