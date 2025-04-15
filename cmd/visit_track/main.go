package main

import (
	"log"
	"net/http"

	visitHttp "github.com/edujudici/visit-track/internal/visit/infra/http"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/visit", visitHttp.VisitHandler)
	r.Get("/visit/total", visitHttp.TotalHandler)

	log.Println("Servidor rodando em http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
