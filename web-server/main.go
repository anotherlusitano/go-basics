package main

import (
	"fmt"
	"net/http"
)

const addr string = ":3333"

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from a Go program"))
	})

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Error: couldn't open the server")
	}
}
