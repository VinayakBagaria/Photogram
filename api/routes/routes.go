package routes

import (
	"net/http"

	"github.com/VinayakBagaria/go-cat-pictures/api/middlewares"
	"github.com/gorilla/mux"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

func Install(router *mux.Router, routeList []*Route) {
	for _, route := range routeList {
		router.Handle(route.Path, middlewares.LogRequests(route.Handler)).Methods(route.Method)
	}
}
