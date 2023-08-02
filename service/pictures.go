package service

import (
	"github.com/VinayakBagaria/go-cat-pictures/db"
	"github.com/VinayakBagaria/go-cat-pictures/dto"
	"github.com/VinayakBagaria/go-cat-pictures/repository"
)

type PicturesService interface {
	List() ([]*db.Picture, error)
	Get(int) (*db.Picture, error)
	Create(dto.CreatePictureInput) (*db.Picture, error)
	Delete(int) error
	Update(int, dto.UpdatePictureInput) (*db.Picture, error)
}

type picturesService struct {
	repository repository.PicturesRepository
}

func NewPicturesService(picturesRepository repository.PicturesRepository) PicturesService {
	return &picturesService{repository: picturesRepository}
}

func (s *picturesService) List() ([]*db.Picture, error) {
	pictures, err := s.repository.GetAll()
	return pictures, err
}

func (s *picturesService) Get(id int) (*db.Picture, error) {
	picture, err := s.repository.GetById(id)
	return picture, err
}

func (s *picturesService) Create(pictureInput dto.CreatePictureInput) (*db.Picture, error) {
	picture, err := s.repository.Create(pictureInput)
	return picture, err
}

func (s *picturesService) Delete(id int) error {
	err := s.repository.Delete(id)
	return err
}

func (s *picturesService) Update(id int, pictureInput dto.UpdatePictureInput) (*db.Picture, error) {
	picture, err := s.repository.Update(id, pictureInput)
	return picture, err
}
