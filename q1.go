package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Q")
}

func main() {
	http.Handle("/", http.HandlerFunc(handler))

	http.ListenAndServe(":3000", nil)
}
