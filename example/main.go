package main

import (
	"io"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", helloHandler)
	srv := http.Server{
		Addr:    ":8080",
		Handler: HTTPLogger(serveMux),
	}
	log.Fatal(srv.ListenAndServe())
}
