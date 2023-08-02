package service

type CreatePictureInput struct {
	Name string `json:"name" binding:"required"`
	Url  string `json:"url" binding:"required"`
}
