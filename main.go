package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := NewRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
