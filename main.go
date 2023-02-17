package main

import (
	"framework"
	"log"
	"net/http"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)

	server := &http.Server{
		Handler: core,
		Addr:    ":8080",
	}

	log.Fatal(server.ListenAndServe())
}
