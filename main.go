package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})
	err := http.ListenAndServe("localhost:8008", mux)
	if err != nil {
		fmt.Println("failed to start server")
	}
}
