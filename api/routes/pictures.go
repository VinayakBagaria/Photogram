package routes

import (
	"net/http"

	"github.com/VinayakBagaria/go-cat-pictures/api/resthandlers"
)

func NewPicturesRoutes(handlers resthandlers.PicturesHandlers) []*Route {
	return []*Route{
		{Path: "/", Method: http.MethodGet, Handler: handlers.GetPictures},
	}
}
