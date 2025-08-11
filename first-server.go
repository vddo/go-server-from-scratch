package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hallo Binh, hier ist Duc, lol.\n")
}

func loveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Liebe, Liebe, Liebe\n")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/love/", loveHandler)
	server := http.Server{
		Addr: ":" + port,
	}
	server.ListenAndServe()
}
