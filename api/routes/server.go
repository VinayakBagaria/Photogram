package routes

import (
	"net/http"

	"github.com/VinayakBagaria/go-cat-pictures/api/resthandlers"
)

func NewServerRouteList(handlers resthandlers.ServerHandler) []*Route {
	return []*Route{
		{Path: "/healthcheck", Method: http.MethodGet, Handler: handlers.HealthCheck},
	}
}
