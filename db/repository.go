package db

import (
	"encoding/json"
	"fmt"

	"github.com/VinayakBagaria/photogram/dto"
	"gorm.io/gorm"
)

type PicturesRepository interface {
	Create(*dto.PictureRequest) (*Picture, error)
	Update(int, *dto.PictureRequest) (*Picture, error)
	Delete(id int) error
	GetAll(int, int) ([]*Picture, int64, error)
	GetById(int) (*Picture, error)
}

type picturesRepository struct {
	db *gorm.DB
}

func NewPicturesRepository(dbHandler *gorm.DB) PicturesRepository {
	return &picturesRepository{db: dbHandler}
}

func (p *picturesRepository) Create(request *dto.PictureRequest) (*Picture, error) {
	picture := Picture{
		Name:        request.Name,
		Destination: request.Destination,
		Height:      request.Height,
		Width:       request.Width,
		Size:        request.Size,
		ContentType: request.ContentType,
	}
	p.db.Create(&picture)
	return &picture, nil
}

func (p *picturesRepository) Update(id int, request *dto.PictureRequest) (*Picture, error) {
	var pictureToUpdate *Picture

	if err := p.db.Where("id = ? AND deleted = ?", id, false).First(&pictureToUpdate).Error; err != nil {
		return nil, err
	}

	marshalledBytes, _ := json.Marshal(request)
	requestMap := make(map[string]interface{})
	json.Unmarshal(marshalledBytes, &requestMap)

	result := p.db.Model(&pictureToUpdate).Where("id = ? AND deleted = ?", id, false).Updates(requestMap)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("record with id: %d not found", id)
	}

	fmt.Println("updating")
	fmt.Println(pictureToUpdate)

	return pictureToUpdate, nil
}

func (p *picturesRepository) Delete(id int) error {
	result := p.db.Where("id = ? AND deleted = ?", id, false).Updates(Picture{Deleted: true})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("record with id: %d not found", id)
	}

	return nil
}

func (p *picturesRepository) GetAll(limit, page int) ([]*Picture, int64, error) {
	var pictures []*Picture
	p.db.Where("deleted = ?", false).Order("updated_on desc").Limit(limit).Offset(limit * (page - 1)).Find(&pictures)
	var totalCount int64
	p.db.Model(&Picture{}).Where("deleted = ?", false).Count(&totalCount)
	return pictures, totalCount, nil
}

func (p *picturesRepository) GetById(id int) (*Picture, error) {
	var picture *Picture

	if err := p.db.Where("id = ? AND deleted = ?", id, false).First(&picture).Error; err != nil {
		return nil, err
	}

	return picture, nil
}
