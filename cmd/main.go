package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Notes API")
	})
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
