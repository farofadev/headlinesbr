package routes

import (
	"github.com/farofadev/headlinesbr/handlers"
	"github.com/julienschmidt/httprouter"
)

func Setup(router *httprouter.Router) *httprouter.Router {
	router.GET("/", handlers.IndexHandler)
	router.GET("/portals", handlers.PortalsIndexHandler)
	router.GET("/headlines", handlers.HeadlinesHandler)

	return router
}
