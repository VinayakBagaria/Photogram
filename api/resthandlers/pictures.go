package resthandlers

import (
	"net/http"

	"github.com/VinayakBagaria/go-cat-pictures/service"
)

type PicturesHandlers interface {
	GetPictures(w http.ResponseWriter, r *http.Request)
}

type picturesHandlers struct {
	picturesService service.PicturesService
}

func NewPicturesHandlers(picturesService service.PicturesService) PicturesHandlers {
	return &picturesHandlers{picturesService}
}

func (h *picturesHandlers) GetPictures(w http.ResponseWriter, r *http.Request) {}
