package dto

import "github.com/gin-gonic/gin"

type PictureRequest struct {
	Name        string
	Destination string
	Height      int32
	Width       int32
	Size        int32
	ContentType string
}

type PictureResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Height      int32  `json:"height"`
	Width       int32  `json:"width"`
	Size        string `json:"size"`
	ContentType string `json:"content_type"`
}

type ListPicturesResponse struct {
	Pictures []*PictureResponse `json:"pictures"`
}

type InvalidPictureFileError struct {
	StatusCode int
	Error      error
	Data       gin.H
}
