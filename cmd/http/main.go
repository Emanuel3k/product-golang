package main

import (
	"github.com/emanuel3k/product-golang/cmd/http/routes"
	"github.com/emanuel3k/product-golang/storage/postgres"
	"log"
	"net/http"
)

func main() {
	if _, err := postgres.Config(); err != nil {
		log.Fatal(err)
	}
	log.Println("Database Connection Successful")

	r := routes.NewRouter()

	var port = ":8080"
	log.Println("Server started on http://localhost:8080 🚀")
	if err := http.ListenAndServe(port, r.MapRoutes()); err != nil {
		log.Fatal(err)
	}
}
