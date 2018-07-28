package main

import (
	"fmt"
	"net/http"
)

type firstHandler struct {
	Sample string
}

func (f firstHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(f.Sample)
	w.Write([]byte(f.Sample))
}

func main() {
	http.Handle("/first", firstHandler{Sample: "This is second test"})
	http.ListenAndServe("0.0.0.0:8080", nil)
}
