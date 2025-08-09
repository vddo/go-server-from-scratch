package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    ":8000",
		Handler: nil,
	}
	server.ListenAndServe()
}
