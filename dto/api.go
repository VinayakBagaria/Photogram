package dto

import (
	"time"

	"github.com/gin-gonic/gin"
)

type PictureRequest struct {
	Name        string
	Destination string
	Height      int32
	Width       int32
	Size        int32
	ContentType string
}

type InvalidPictureFileError struct {
	StatusCode int
	Error      error
	Data       gin.H
}

type PictureResponse struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Url         string    `json:"url"`
	Height      int32     `json:"height"`
	Width       int32     `json:"width"`
	Size        string    `json:"size"`
	ContentType string    `json:"content_type"`
	CreatedOn   time.Time `json:"created_on"`
	UpdatedOn   time.Time `json:"updated_on"`
}

type ListPicturesResponse struct {
	Pictures   []*PictureResponse `json:"pictures"`
	Count      int                `json:"count"`
	TotalPages int                `json:"total_pages"`
}

type SinglePictureResponse struct {
	Data *PictureResponse `json:"data"`
}

type StringResponse struct {
	Message string `json:"message"`
}

type GeneralErrorResponse struct {
	Error string         `json:"error"`
	Meta  map[string]any `json:"meta,omitempty"`
}
