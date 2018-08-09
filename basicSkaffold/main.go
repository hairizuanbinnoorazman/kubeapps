package main

import (
	"log"
	"net/http"
)

type HelloWorld struct{}

func (l HelloWorld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello World http endpoint reached")
	w.Write([]byte("This is a sample"))
}

func main() {
	log.Println("Server started")
	http.Handle("/", HelloWorld{})
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
