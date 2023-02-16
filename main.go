package main

import (
	"framework"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Handler: framework.NewCore(),
		Addr:    ":8080",
	}

	log.Fatal(server.ListenAndServe())
}
