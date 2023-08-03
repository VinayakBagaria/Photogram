package dto

type CreatePictureRequest struct {
	Name        string `form:"name" binding:"required"`
	Destination string `form:"destination" binding:"required"`
}

type UpdatePictureRequest struct {
	Name string `form:"name"`
	Url  string `form:"url"`
}

type PictureResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}
