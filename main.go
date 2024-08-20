package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
)

func main() {
	server := &http.Server{
		Addr: ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("An error occurred: ", err)
	}

	// Routing will be done with if checks in this method
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}