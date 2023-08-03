package dto

type PictureRequest struct {
	Name        string
	Destination string
	Height      int
	Width       int
	Size        int64
	ContentType string
}

type PictureResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ListPicturesResponse struct {
	Pictures []*PictureResponse `json:"pictures"`
}

type InvalidPictureFileError struct {
	StatusCode int
	Error      error
}
