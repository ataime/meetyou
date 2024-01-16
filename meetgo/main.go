package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	}))
}
