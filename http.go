package main

import (
	"net/http"
)

type myHandler struct {
	s string
}

func (myH *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("my handler: " + myH.s))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Main)

	myH := &myHandler{
		s: "test my handler",
	}
	mux.Handle("/my-handler", myH)

	http.ListenAndServe(":3000", mux)
}

func Main(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("main max"))
}
