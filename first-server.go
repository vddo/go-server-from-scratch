package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello, world\n")
}

func main() {
	http.HandleFunc("/", helloHandler)
	server := http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()
}
