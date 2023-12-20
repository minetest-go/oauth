package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting")

	server := &http.Server{Addr: ":8080", Handler: nil}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
