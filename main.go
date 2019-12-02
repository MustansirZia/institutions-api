package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mustansirzia/institutions-api/institutions"
)

func main() {

	addr := ":" + getPort()

	mux := http.NewServeMux()
	mux.HandleFunc("/institutions", institutions.HandleHTTP)

	fmt.Printf("About to listen on %s!\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
	}
	return port
}
