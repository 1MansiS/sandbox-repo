package main

import (
	"log"
	"net/http"

	"sandbox-repo/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Hello)
	http.HandleFunc("/files", handlers.ServeFile)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
