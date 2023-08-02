package repository

import "github.com/VinayakBagaria/go-cat-pictures/db"

type PicturesRepository interface {
	Save(*db.Picture) error
	GetById(string) (*db.Picture, error)
	GetAll() ([]*db.Picture, error)
	Update(*db.Picture) error
}

type picturesRepository struct {
}

func NewPicturesRepository() PicturesRepository {
	return &picturesRepository{}
}

func (p *picturesRepository) Save(picture *db.Picture) error {
	return nil
}

func (p *picturesRepository) GetById(id string) (*db.Picture, error) {
	return nil, nil
}

func (p *picturesRepository) GetAll() ([]*db.Picture, error) {
	return []*db.Picture{}, nil
}

func (p *picturesRepository) Update(picture *db.Picture) error {
	return nil
}
