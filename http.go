package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/")

	fmt.Printf("%#v\n", mux)
}

func Main(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("main"))
}
