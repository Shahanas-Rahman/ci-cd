package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from the Go Backend!")
	})
	http.ListenAndServe(":8081", nil)
}
