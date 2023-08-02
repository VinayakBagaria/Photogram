package service

import (
	"github.com/VinayakBagaria/go-cat-pictures/db"
	"github.com/VinayakBagaria/go-cat-pictures/repository"
)

type PicturesService interface {
	ListPictures() ([]*db.Picture, error)
	CreatePicture(CreatePictureInput) (db.Picture, error)
}

type picturesService struct {
	picturesRepository repository.PicturesRepository
}

func NewPicturesService(picturesRepository repository.PicturesRepository) PicturesService {
	return &picturesService{picturesRepository}
}

func (s *picturesService) ListPictures() ([]*db.Picture, error) {
	pictures, err := s.picturesRepository.GetAll()
	return pictures, err
}

func (s *picturesService) CreatePicture(pictureInput CreatePictureInput) (db.Picture, error) {
	picture := db.Picture{Name: pictureInput.Name, Url: pictureInput.Url}
	err := s.picturesRepository.Create(&picture)
	return picture, err
}
