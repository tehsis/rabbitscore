package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tehsis/rabbitscore/middlewares"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = SetContentType(handler, "application/json")
		handler = middlewares.Logger(handler, route.Name)
		handler = middlewares.Authorize(handler)

		router.Path(route.Pattern).Methods(route.Method).Handler(handler)
	}

	return router
}
