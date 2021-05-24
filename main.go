package main

import (
	"log"
	"net/http"
	"os"

	"github.com/farofadev/headlinesbr/routes"
	"github.com/farofadev/headlinesbr/scheduler"
	"github.com/julienschmidt/httprouter"
)

func main() {
	go scheduler.Run()

	port := os.Getenv("SERVER_PORT")

	if port == "" {
		port = "8080"
	}

	router := routes.Setup(httprouter.New())
	log.Fatal(http.ListenAndServe(":"+port, router))
}
