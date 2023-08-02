package service

import (
	"github.com/VinayakBagaria/go-cat-pictures/db"
	"github.com/VinayakBagaria/go-cat-pictures/repository"
)

type PicturesService interface {
	ListPictures() ([]*db.Picture, error)
	GetPicture(int) (*db.Picture, error)
	CreatePicture(CreatePictureInput) (db.Picture, error)
	Delete(int) error
}

type picturesService struct {
	repository repository.PicturesRepository
}

func NewPicturesService(picturesRepository repository.PicturesRepository) PicturesService {
	return &picturesService{repository: picturesRepository}
}

func (s *picturesService) ListPictures() ([]*db.Picture, error) {
	pictures, err := s.repository.GetAll()
	return pictures, err
}

func (s *picturesService) GetPicture(id int) (*db.Picture, error) {
	picture, err := s.repository.GetById(id)
	return picture, err
}

func (s *picturesService) CreatePicture(pictureInput CreatePictureInput) (db.Picture, error) {
	picture := db.Picture{Name: pictureInput.Name, Url: pictureInput.Url}
	err := s.repository.Create(&picture)
	return picture, err
}

func (s *picturesService) Delete(id int) error {
	err := s.repository.Delete(id)
	return err
}
