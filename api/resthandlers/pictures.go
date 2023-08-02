package resthandlers

import (
	"net/http"

	"github.com/VinayakBagaria/go-cat-pictures/api/restutil"
	"github.com/VinayakBagaria/go-cat-pictures/service"
	"github.com/gin-gonic/gin"
)

type PicturesHandler interface {
	GetPictures(c *gin.Context)
	CreatePictures(c *gin.Context)
}

type picturesHandler struct {
	picturesService service.PicturesService
}

func NewPicturesHandlers(picturesService service.PicturesService) PicturesHandler {
	return &picturesHandler{picturesService}
}

func (h *picturesHandler) GetPictures(c *gin.Context) {
	pictures, err := h.picturesService.ListPictures()
	if err != nil {
		restutil.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	restutil.WriteAsJson(c, http.StatusOK, gin.H{"pictures": pictures})
}

func (h *picturesHandler) CreatePictures(c *gin.Context) {
	var input service.CreatePictureInput
	if err := c.ShouldBindJSON(&input); err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err)
		return
	}

	picture, err := h.picturesService.CreatePicture(input)
	if err != nil {
		restutil.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	restutil.WriteAsJson(c, http.StatusOK, gin.H{"data": picture})
}
