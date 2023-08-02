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
	CreatePicture(*gin.Context)
	DeletePicture(*gin.Context)
	UpdatePicture(*gin.Context)
}

type picturesHandler struct {
	svc service.PicturesService
}

func NewPicturesHandlers(picturesService service.PicturesService) PicturesHandler {
	return &picturesHandler{svc: picturesService}
}

func (h *picturesHandler) ListPictures(c *gin.Context) {
	pictures, err := h.svc.List()
	if err != nil {
		restutil.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	restutil.WriteAsJson(c, http.StatusOK, gin.H{"pictures": pictures})
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

func (h *picturesHandler) CreatePicture(c *gin.Context) {
	var input dto.CreatePictureInput
	if err := c.ShouldBindJSON(&input); err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err)
		return
	}

	picture, err := h.svc.Create(input)
	if err != nil {
		restutil.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	restutil.WriteAsJson(c, http.StatusOK, gin.H{"data": picture})
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

	var input dto.UpdatePictureInput
	if err := c.ShouldBindJSON(&input); err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err)
		return
	}

	picture, err := h.svc.Update(id, input)
	if err != nil {
		restutil.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	restutil.WriteAsJson(c, http.StatusOK, gin.H{"data": picture})
}
