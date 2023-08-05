package repository

import (
	"encoding/json"
	"fmt"

	"github.com/VinayakBagaria/go-cat-pictures/db"
	"github.com/VinayakBagaria/go-cat-pictures/dto"
	"gorm.io/gorm"
)

type PicturesRepository interface {
	Create(*dto.PictureRequest) (*db.Picture, error)
	Update(int, *dto.PictureRequest) (*db.Picture, error)
	Delete(id int) error
	GetAll(int, int) ([]*db.Picture, int64, error)
	GetById(int) (*db.Picture, error)
}

type picturesRepository struct {
	db *gorm.DB
}

func NewPicturesRepository(dbHandler *gorm.DB) PicturesRepository {
	return &picturesRepository{db: dbHandler}
}

func (p *picturesRepository) Create(request *dto.PictureRequest) (*db.Picture, error) {
	picture := db.Picture{
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

func (p *picturesRepository) Update(id int, request *dto.PictureRequest) (*db.Picture, error) {
	var pictureToUpdate *db.Picture
	if err := p.db.Where("id = ?", id).First(&pictureToUpdate).Error; err != nil {
		return nil, err
	}

	marshalledBytes, _ := json.Marshal(request)
	requestMap := make(map[string]interface{})
	json.Unmarshal(marshalledBytes, &requestMap)

	p.db.Model(&pictureToUpdate).Updates(requestMap)
	return pictureToUpdate, nil
}

func (p *picturesRepository) Delete(id int) error {
	result := p.db.Where("id = ?", id).Updates(db.Picture{Deleted: true})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("record with id: %d not found", id)
	}

	return nil
}

func (p *picturesRepository) GetAll(limit, page int) ([]*db.Picture, int64, error) {
	var pictures []*db.Picture
	p.db.Where("deleted = ?", false).Order("updated_on desc").Limit(limit).Offset(limit * (page - 1)).Find(&pictures)
	var totalCount int64
	p.db.Model(&db.Picture{}).Where("deleted = ?", false).Count(&totalCount)
	return pictures, totalCount, nil
}

func (p *picturesRepository) GetById(id int) (*db.Picture, error) {
	var picture *db.Picture

	if err := p.db.Where("id = ? AND deleted = ?", id, false).First(&picture).Error; err != nil {
		return nil, err
	}

	return picture, nil
}
