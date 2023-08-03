package resthandlers

import (
	"net/http"
	"strconv"

	"github.com/VinayakBagaria/go-cat-pictures/api/restutil"
	"github.com/VinayakBagaria/go-cat-pictures/dto"
	"github.com/VinayakBagaria/go-cat-pictures/service"
	"github.com/gin-gonic/gin"
)

type PicturesHandler interface {
	ListPictures(*gin.Context)
	GetPicture(*gin.Context)
	GetPictureFile(*gin.Context)
	CreatePicture(*gin.Context)
	DeletePicture(*gin.Context)
	UpdatePicture(*gin.Context)
}

type picturesHandler struct {
	svc service.PicturesService
}

func NewPicturesHandler(picturesService service.PicturesService) PicturesHandler {
	return &picturesHandler{svc: picturesService}
}

// List of pictures along with its metadata
// @Summary list of pictures
// @Success 200 {object} dto.ListPicturesResponse
// @Failure 500 {object} error
// @Router / [get]
func (h *picturesHandler) ListPictures(c *gin.Context) {
	pictures, err := h.svc.List()
	if err != nil {
		restutil.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	restutil.WriteAsJson(c, http.StatusOK, dto.ListPicturesResponse{Pictures: pictures})
}

func (h *picturesHandler) GetPicture(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err)
		return
	}

	picture, err := h.svc.Get(id)
	if err != nil {
		restutil.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	restutil.WriteAsJson(c, http.StatusOK, gin.H{"data": picture})
}

func (h *picturesHandler) GetPictureFile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err)
		return
	}

	pictureDestination, err := h.svc.GetFile(id)
	if err != nil {
		restutil.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	http.ServeFile(c.Writer, c.Request, "./images/"+pictureDestination)
}

func (h *picturesHandler) CreatePicture(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err)
		return
	}

	response, err := h.svc.Create(file)
	if err != nil {
		restutil.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	restutil.WriteAsJson(c, http.StatusOK, gin.H{"data": response})
}

func (h *picturesHandler) DeletePicture(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err)
		return
	}

	if err := h.svc.Delete(id); err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err)
		return
	}

	restutil.WriteAsJson(c, http.StatusOK, gin.H{"data": true})
}

func (h *picturesHandler) UpdatePicture(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err)
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err)
		return
	}

	picture, err := h.svc.Update(id, file)
	if err != nil {
		restutil.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	restutil.WriteAsJson(c, http.StatusOK, gin.H{"data": picture})
}
