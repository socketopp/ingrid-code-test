package main

import (
	"fmt"
	"log"
	"net/http"

	"home/server"
)

func main() {
	fmt.Println("Building REST API")
	mux := server.NewServeMux()
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
