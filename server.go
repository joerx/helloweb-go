package main

import (
	"log"
	"net/http"

	"github.com/joerx/helloweb/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HandleIndex)
	log.Printf("Starting")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
