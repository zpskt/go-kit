package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world received a request.")
	version := os.Getenv("VERSION")
	if version == "" {
		version = "v1"
	}
	log.Println(version)
	fmt.Fprintf(w, "Hello world %s\n", version)
}
func main() {
	log.Print("Hello world sample started.")
	http.HandleFunc("/api/hello", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
