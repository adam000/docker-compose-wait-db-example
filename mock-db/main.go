package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	time.Sleep(10)
	http.HandleFunc("/db", dbQuery)
	http.ListenAndServe(":1212", nil)
}

func dbQuery(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "database result\n")
}
