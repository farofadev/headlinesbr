package main

import (
	"log"
	"net/http"

	"github.com/farofadev/headlinesbr/routes"
	"github.com/farofadev/headlinesbr/scheduler"
	"github.com/julienschmidt/httprouter"
)

func main() {
	go scheduler.Run()

	router := routes.Setup(httprouter.New())

	log.Fatal(http.ListenAndServe(":8080", router))
}
