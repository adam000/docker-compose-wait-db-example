package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", hello)
	mux.HandleFunc("GET /healthcheck", func(w http.ResponseWriter, req *http.Request) {})
	http.ListenAndServe(":8080", mux)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "If you're seeing this, then migrations should have successfully completed\n")
}
