package main

import (
	"framework"
	"framework/middlewares"
	"log"
	"net/http"
	"time"
)

func main() {
	core := framework.NewCore()

	core.Use(middlewares.Recovery())
	core.Use(middlewares.Cost())
	core.Use(middlewares.Timeout(1 * time.Second))

	registerRouter(core)

	server := &http.Server{
		Handler: core,
		Addr:    ":8080",
	}

	log.Fatal(server.ListenAndServe())
}
