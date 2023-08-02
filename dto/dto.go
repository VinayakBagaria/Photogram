package dto

type CreatePictureInput struct {
	Name string `json:"name" binding:"required"`
	Url  string `json:"url" binding:"required"`
}

type UpdatePictureInput struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
