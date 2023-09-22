package db

import (
	"fmt"
	"time"

	"github.com/VinayakBagaria/photogram/config"
	"github.com/VinayakBagaria/photogram/dto"
)

type Picture struct {
	ID        uint  `json:"id" gorm:"primary_key"`
	CreatedOn int64 `json:"created_on" gorm:"autoCreateTime:milli"`
	UpdatedOn int64 `json:"updated_on" gorm:"autoUpdateTime:milli"`
	Deleted   bool  `json:"deleted" gorm:"default:false"`

	Name        string `json:"name"`
	Destination string `json:"destination"`
	Height      int32  `json:"height"`
	Width       int32  `json:"width"`
	Size        int32  `json:"size"`
	ContentType string `json:"content_type"`
}


func (p *Picture) ToPictureResponse() *dto.PictureResponse {
	return &dto.PictureResponse{
		Id:          p.ID,
		Name:        p.Name,
		Url:         fmt.Sprintf("%s/picture/%d/image", config.GetConfigValue("server.host"), p.ID),
		Height:      p.Height,
		Width:       p.Width,
		Size:        fmt.Sprintf("%.2f KB", float64(p.Size)/1024),
		ContentType: p.ContentType,
		CreatedOn:   time.UnixMilli(p.CreatedOn),
		UpdatedOn:   time.UnixMilli(p.UpdatedOn),
	}
}
