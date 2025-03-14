package main

import (
	"fmt"
	"github.com/emanuel3k/product-golang/cmd/http/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.NewRouter()

	var port = ":8080"
	fmt.Println("Server started on http://localhost:8080 🚀")
	if err := http.ListenAndServe(port, r.MapRoutes()); err != nil {
		log.Fatal(err)
	}
}
