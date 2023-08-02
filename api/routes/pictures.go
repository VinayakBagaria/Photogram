package routes

import (
	"net/http"

	"github.com/VinayakBagaria/go-cat-pictures/api/resthandlers"
)

func NewPicturesRoutes(handlers resthandlers.PicturesHandler) []*Route {
	return []*Route{
		{Path: "/", Method: http.MethodGet, Handler: handlers.GetPictures},
		{Path: "/", Method: http.MethodPost, Handler: handlers.CreatePictures},
	}
}
