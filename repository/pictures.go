package repository

import (
	"github.com/VinayakBagaria/go-cat-pictures/db"
	"gorm.io/gorm"
)

type PicturesRepository interface {
	Create(*db.Picture) error
	GetById(string) (*db.Picture, error)
	GetAll() ([]*db.Picture, error)
	Update(*db.Picture) error
	Delete(id string) error
}

type picturesRepository struct {
	db *gorm.DB
}

func NewPicturesRepository(dbHandler *gorm.DB) PicturesRepository {
	return &picturesRepository{db: dbHandler}
}

func (p *picturesRepository) Create(picture *db.Picture) error {
	p.db.Create(&picture)
	return nil
}

func (p *picturesRepository) GetById(id string) (*db.Picture, error) {
	return nil, nil
}

func (p *picturesRepository) GetAll() ([]*db.Picture, error) {
	var pictures []*db.Picture
	p.db.Find(&pictures)
	return pictures, nil
}

func (p *picturesRepository) Update(picture *db.Picture) error {
	return nil
}

func (p *picturesRepository) Delete(id string) error {
	var picture db.Picture

	if err := p.db.Where("?id=", id).First(&picture).Error; err != nil {
		return err
	}

	p.db.Delete(&picture)
	return nil
}
