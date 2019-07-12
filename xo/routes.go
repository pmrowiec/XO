package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GameGetState",
		"POST",
		"/game/{gameId}",
		GameGetState,
	},
	Route{
		"GameMakeMove",
		"POST",
		"/game/move",
		GameMakeMove,
	},
	Route{
		"GameCreate",
		"POST",
		"/game",
		GameCreate,
	},
}
