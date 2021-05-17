package routes

import (
	"github.com/julienschmidt/httprouter"
)

func Setup(router *httprouter.Router) *httprouter.Router {
	router.GET("/", IndexHandler)
	router.GET("/headlines", HeadlinesHandler)

	return router
}
