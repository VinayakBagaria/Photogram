package resthandlers

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/VinayakBagaria/go-cat-pictures/api/restutil"
	"github.com/VinayakBagaria/go-cat-pictures/dto"
	"github.com/VinayakBagaria/go-cat-pictures/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension

	destination := "./images/" + newFileName

	if err := c.SaveUploadedFile(file, destination); err != nil {
		restutil.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	var request dto.CreatePictureRequest
	request.Name = file.Filename
	request.Destination = newFileName
	// if err := c.ShouldBind(&request); err != nil {
	// 	restutil.WriteAsJson(c, http.StatusBadRequest, gin.H{"error_occurred": err.Error()})
	// 	return
	// }

	response, err := h.svc.Create(request)
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

	var request dto.UpdatePictureRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err)
		return
	}

	picture, err := h.svc.Update(id, request)
	if err != nil {
		restutil.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	restutil.WriteAsJson(c, http.StatusOK, gin.H{"data": picture})
}
