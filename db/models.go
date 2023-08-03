package db

import "github.com/VinayakBagaria/go-cat-pictures/dto"

type Picture struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Destination string `json:"destination"`
}

func (p *Picture) ToPictureResponse() dto.PictureResponse {
	return dto.PictureResponse{Id: p.ID, Name: p.Name, Url: "/get-image/" + p.Destination}
}
