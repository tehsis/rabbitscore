package main

import (
	"net/http"

	"github.com/tehsis/rabbitscore/handlers"
)

// Route is a REST resource
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a colleciton of Rest resources
type Routes []Route

var routes = Routes{
	Route{
		"Leaderboard",
		"GET",
		"/leaderboard",
		handlers.LeaderBoardHandler,
	},
	Route{
		"AddRecord",
		"POST",
		"/leaderboard",
		handlers.AddScore,
	},
	Route{
		"Status",
		"GET",
		"/status",
		Status,
	},
}
