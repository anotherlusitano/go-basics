package main

import (
	"fmt"
	"net/http"
)

const addr string = ":3333"

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from a Go program"))
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(addr, server)
	if err != nil {
		fmt.Println("Error: couldn't open the server")
	}
}
