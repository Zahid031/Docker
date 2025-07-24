package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Go")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Starting server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
