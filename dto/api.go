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
	Pictures []*PictureResponse `json:"pictures"`
	Count    int64              `json:"count"`
}

type InvalidPictureFileError struct {
	StatusCode int
	Error      error
	Data       gin.H
}

type PaginationResponse struct {
	Count int
	Data  gin.H
}
