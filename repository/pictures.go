package repository

import (
	"github.com/VinayakBagaria/go-cat-pictures/db"
	"github.com/VinayakBagaria/go-cat-pictures/dto"
	"gorm.io/gorm"
)

type PicturesRepository interface {
	Create(dto.CreatePictureInput) (*db.Picture, error)
	GetById(int) (*db.Picture, error)
	GetAll() ([]*db.Picture, error)
	Delete(id int) error
	Update(int, dto.UpdatePictureInput) (*db.Picture, error)
}

type picturesRepository struct {
	db *gorm.DB
}

func NewPicturesRepository(dbHandler *gorm.DB) PicturesRepository {
	return &picturesRepository{db: dbHandler}
}

func (p *picturesRepository) Create(pictureInput dto.CreatePictureInput) (*db.Picture, error) {
	picture := db.Picture{Name: pictureInput.Name, Url: pictureInput.Url}
	p.db.Create(&picture)
	return &picture, nil
}

func (p *picturesRepository) GetById(id int) (*db.Picture, error) {
	var picture *db.Picture

	if err := p.db.Where("id=?", id).First(&picture).Error; err != nil {
		return nil, err
	}

	return picture, nil
}

func (p *picturesRepository) GetAll() ([]*db.Picture, error) {
	var pictures []*db.Picture
	p.db.Find(&pictures)
	return pictures, nil
}

func (p *picturesRepository) Delete(id int) error {
	var picture db.Picture
	p.db.Where("id=?", id).Delete(&picture)
	return nil
}

func (p *picturesRepository) Update(id int, pictureInput dto.UpdatePictureInput) (*db.Picture, error) {
	var pictureToUpdate *db.Picture
	if err := p.db.Where("id=?", id).First(&pictureToUpdate).Error; err != nil {
		return nil, err
	}

	p.db.Model(&pictureToUpdate).Updates(&pictureInput)
	return pictureToUpdate, nil
}
