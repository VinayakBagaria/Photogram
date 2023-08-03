package db

import (
	"fmt"

	"github.com/VinayakBagaria/go-cat-pictures/dto"
)

type Picture struct {
	ID        uint  `json:"id" gorm:"primary_key"`
	CreatedOn int64 `json:"created_on" gorm:"autoCreateTime:milli"`
	UpdatedOn int64 `json:"updated_on" gorm:"autoUpdateTime:milli"`
	Deleted   bool  `json:"deleted" gorm:"default:false"`

	Name        string `json:"name"`
	Destination string `json:"destination"`
	Height      int    `json:"height"`
	Width       int    `json:"width"`
	Size        int64  `json:"size"`
	ContentType string `json:"content_type"`
}

func (p *Picture) ToPictureResponse() *dto.PictureResponse {
	return &dto.PictureResponse{Id: p.ID, Name: p.Name, Url: fmt.Sprintf("/picture/%d/image", p.ID)}
}
