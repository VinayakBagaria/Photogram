package resthandlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/VinayakBagaria/go-cat-pictures/api/restutil"
	"github.com/VinayakBagaria/go-cat-pictures/dto"
	"github.com/VinayakBagaria/go-cat-pictures/service"
	"github.com/gin-gonic/gin"
)

type PicturesHandler interface {
	CreatePicture(*gin.Context)
	UpdatePicture(*gin.Context)
	ListPictures(*gin.Context)
	GetPicture(*gin.Context)
	GetPictureFile(*gin.Context)
	DeletePicture(*gin.Context)
}

type picturesHandler struct {
	svc service.PicturesService
}

func NewPicturesHandler(picturesService service.PicturesService) PicturesHandler {
	return &picturesHandler{svc: picturesService}
}

func (h *picturesHandler) CreatePicture(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err, nil)
		return
	}

	response, createError := h.svc.Create(file)
	if createError != nil {
		restutil.WriteError(c, createError.StatusCode, createError.Error, createError.Data)
		return
	}

	restutil.WriteAsJson(c, http.StatusOK, gin.H{"data": response})
}

func (h *picturesHandler) UpdatePicture(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err, nil)
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err, nil)
		return
	}

	response, updatedError := h.svc.Update(id, file)
	if err != nil {
		restutil.WriteError(c, updatedError.StatusCode, updatedError.Error, nil)
		return
	}

	restutil.WriteAsJson(c, http.StatusOK, gin.H{"data": response})
}

// List of pictures
// @Summary list of pictures
// @Description List of pictures along with its metadata
// @Param page query string false "page number starting from 1" Format(number)
// @Success 200 {object} dto.ListPicturesResponse
// @Failure 500 {object} error
// @Router / [get]
func (h *picturesHandler) ListPictures(c *gin.Context) {
	pageSize := 10
	page := c.Query("page")
	if page == "" {
		page = "1"
	}

	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err, nil)
		return
	}

	if pageNumber < 1 {
		restutil.WriteError(c, http.StatusBadRequest, errors.New("page can't be less than 1"), nil)
		return
	}

	pictures, totalCount, err := h.svc.List(pageSize, pageNumber)
	if err != nil {
		restutil.WriteError(c, http.StatusInternalServerError, err, nil)
		return
	}

	totalPages := totalCount / pageSize
	if (totalCount % pageSize) > 0 {
		totalPages += 1
	}

	restutil.WriteAsJson(c, http.StatusOK, dto.ListPicturesResponse{
		Pictures:   pictures,
		Count:      totalCount,
		TotalPages: totalPages,
	})
}

// Get a image
// @Summary get a image
// @Description Get a specified image file by its ID
// @Param id path number true "Image Id"
// @Success 200 {file} octet-stream
// @Failure 500 {object} error
// @Router /picture/{id}/image [get]
func (h *picturesHandler) GetPictureFile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err, nil)
		return
	}

	pictureDestination, err := h.svc.GetFile(id)
	if err != nil {
		restutil.WriteError(c, http.StatusInternalServerError, err, nil)
		return
	}

	http.ServeFile(c.Writer, c.Request, pictureDestination)
}

func (h *picturesHandler) GetPicture(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err, nil)
		return
	}

	picture, err := h.svc.Get(id)
	if err != nil {
		restutil.WriteError(c, http.StatusInternalServerError, err, nil)
		return
	}

	restutil.WriteAsJson(c, http.StatusOK, gin.H{"data": picture})
}

func (h *picturesHandler) DeletePicture(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err, nil)
		return
	}

	if err := h.svc.Delete(id); err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err, nil)
		return
	}

	restutil.WriteAsJson(c, http.StatusOK, gin.H{"data": true})
}
