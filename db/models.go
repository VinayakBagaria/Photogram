package db

import (
	"fmt"

	"github.com/VinayakBagaria/go-cat-pictures/dto"
)

type BaseModel struct {
	CreatedOn string `json:"created_on"`
	UpdatedOn string `json:"updated_on"`
	Deleted   bool   `json:"deleted"`
}

type Picture struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Destination string `json:"destination"`
	BaseModel
}

func (p *Picture) ToPictureResponse() dto.PictureResponse {
	return dto.PictureResponse{Id: p.ID, Name: p.Name, Url: fmt.Sprintf("/picture/%d/image", p.ID)}
}
