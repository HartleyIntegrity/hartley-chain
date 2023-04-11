package main

import (
	"net/http"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"/tenancies",
		"GET",
		GetAllTenancies,
	},
	Route{
		"/tenancies/{id}",
		"GET",
		GetTenancy,
	},
	Route{
		"/tenancies",
		"POST",
		CreateTenancy,
	},
	Route{
		"/tenancies/{id}",
		"PUT",
		UpdateTenancy,
	},
	Route{
		"/tenancies/{id}",
		"DELETE",
		DeleteTenancy,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Handler(route.Handler)
	}

	return router
}
