package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hallo Binh, hier ist Duc, lol.\n")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", helloHandler)
	server := http.Server{
		Addr: ":" + port,
	}
	server.ListenAndServe()
}
