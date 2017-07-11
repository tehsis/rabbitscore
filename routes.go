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
	Protected   bool
	HandlerFunc http.HandlerFunc
}

// Routes is a colleciton of Rest resources
type Routes []Route

var routes = Routes{
	Route{
		"Leaderboard",
		"GET",
		"/leaderboard",
		false,
		handlers.LeaderBoardHandler,
	},
	Route{
		"Authentication",
		"POST",
		"/login",
		false,
		handlers.AuthenticationHandler,
	},
	Route{
		"AddRecord",
		"POST",
		"/leaderboard",
		true,
		handlers.AddScore,
	},
	Route{
		"Status",
		"GET",
		"/status",
		false,
		Status,
	},
}
