package routes

import (
	"net/http"

	"github.com/VinayakBagaria/go-cat-pictures/api/resthandlers"
)

func NewPicturesRoutes(handlers resthandlers.PicturesHandler) []*Route {
	return []*Route{
		{Path: "/", Method: http.MethodGet, Handler: handlers.ListPictures},
		{Path: "/picture/:id", Method: http.MethodGet, Handler: handlers.GetPicture},
		{Path: "/", Method: http.MethodPost, Handler: handlers.CreatePicture},
		{Path: "/picture/:id", Method: http.MethodDelete, Handler: handlers.DeletePicture},
		{Path: "/picture/:id", Method: http.MethodPut, Handler: handlers.UpdatePicture},
	}
}
