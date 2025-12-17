package main

import (
	"log"
	"net/http"

	"folio/routes"
)

func main() {
	r := routes.RegisterRoutes()

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))	
}