package repository

import "github.com/VinayakBagaria/go-cat-pictures/models"

type PicturesRepository interface {
	Save(*models.Picture) error
	GetById(string) (*models.Picture, error)
	GetAll() ([]*models.Picture, error)
	Update(*models.Picture) error
}

type picturesRepository struct {
}

func NewPicturesRepository() PicturesRepository {
	return &picturesRepository{}
}

func (p *picturesRepository) Save(picture *models.Picture) error {
	return nil
}

func (p *picturesRepository) GetById(id string) (*models.Picture, error) {
	return nil, nil
}

func (p *picturesRepository) GetAll() ([]*models.Picture, error) {
	return []*models.Picture{}, nil
}

func (p *picturesRepository) Update(picture *models.Picture) error {
	return nil
}
