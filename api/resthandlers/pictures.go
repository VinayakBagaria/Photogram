package resthandlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/VinayakBagaria/photogram/api/restutil"
	"github.com/VinayakBagaria/photogram/dto"
	"github.com/VinayakBagaria/photogram/service"
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

// Save an image
// @Summary save an image
// @Description Given a image file, save it & get its computed metadata
// @Accept			multipart/form-data
//
//	@Param			image	formData	file			true	"upload image file"
//
// @Success 201 {object} dto.SinglePictureResponse
// @Failure 400 {object} dto.GeneralErrorResponse
// @Failure 500 {object} dto.GeneralErrorResponse
// @Router / [post]
func (h *picturesHandler) CreatePicture(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err, nil)
		return
	}

	createdPicture, createError := h.svc.Create(file)
	if createError != nil {
		restutil.WriteError(c, createError.StatusCode, createError.Error, createError.Data)
		return
	}

	restutil.WriteAsJson(c, http.StatusCreated, dto.SinglePictureResponse{Data: createdPicture})
}

// Update an image
// @Summary update an image
// @Description Given a image file and an id, update the record & get its computed metadata
// @Accept			multipart/form-data
// @Param id path number true "Image Id"
//
//	@Param			image	formData	file			true	"upload image file"
//
// @Success 202 {object} dto.SinglePictureResponse
// @Failure 400 {object} dto.GeneralErrorResponse
// @Failure 404 {object} dto.GeneralErrorResponse
// @Failure 500 {object} dto.GeneralErrorResponse
// @Router /picture/{id} [put]
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

	pictureResponse, updatedError := h.svc.Update(id, file)
	if updatedError != nil {
		restutil.WriteError(c, updatedError.StatusCode, updatedError.Error, nil)
		return
	}

	restutil.WriteAsJson(c, http.StatusAccepted, dto.SinglePictureResponse{Data: pictureResponse})
}

// List of pictures
// @Summary list of pictures
// @Description List of pictures along with its metadata
// @Param page query number false "page number starting from 1" Format(number)
// @Success 200 {object} dto.ListPicturesResponse
// @Failure 400 {object} dto.GeneralErrorResponse
// @Failure 500 {object} dto.GeneralErrorResponse
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
// @Failure 400 {object} dto.GeneralErrorResponse
// @Failure 404 {object} dto.GeneralErrorResponse
// @Router /picture/{id}/image [get]
func (h *picturesHandler) GetPictureFile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err, nil)
		return
	}

	pictureDestination, err := h.svc.GetFile(id)
	if err != nil {
		restutil.WriteError(c, http.StatusNotFound, err, nil)
		return
	}

	http.ServeFile(c.Writer, c.Request, pictureDestination)
}

// Get a single image data
// @Summary get a single image data
// @Description Get a specified image with its metadata by its ID
// @Param id path number true "Image Id"
// @Success 200 {object} dto.SinglePictureResponse
// @Failure 400 {object} dto.GeneralErrorResponse
// @Failure 404 {object} dto.GeneralErrorResponse
// @Router /picture/{id} [get]
func (h *picturesHandler) GetPicture(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err, nil)
		return
	}

	picture, err := h.svc.Get(id)
	if err != nil {
		restutil.WriteError(c, http.StatusNotFound, err, nil)
		return
	}

	restutil.WriteAsJson(c, http.StatusOK, dto.SinglePictureResponse{Data: picture})
}

// Delete a single image
// @Summary delete a single image
// @Description Delete a specified image along with its metadata by its ID
// @Param id path number true "Image Id"
// @Success 200 {object} dto.StringResponse
// @Failure 404 {object} dto.GeneralErrorResponse
// @Failure 500 {object} dto.GeneralErrorResponse
// @Router /picture/{id} [delete]
func (h *picturesHandler) DeletePicture(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restutil.WriteError(c, http.StatusBadRequest, err, nil)
		return
	}

	if err := h.svc.Delete(id); err != nil {
		restutil.WriteError(c, http.StatusNotFound, err, nil)
		return
	}

	restutil.WriteAsJson(c, http.StatusOK, dto.StringResponse{Message: "Successfully deleted"})
}
