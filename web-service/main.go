package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", hello)
	mux.HandleFunc("GET /healthcheck", func(w http.ResponseWriter, req *http.Request) {})
	log.Println("Web server has started up!")
	http.ListenAndServe(":8080", mux)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "If you're seeing this, then the API service should have successfully started up\n")
}
