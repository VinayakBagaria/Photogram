package service

import (
	"mime/multipart"

	"github.com/VinayakBagaria/go-cat-pictures/dto"
	"github.com/VinayakBagaria/go-cat-pictures/repository"
	"github.com/VinayakBagaria/go-cat-pictures/storage"
)

type PicturesService interface {
	List() ([]dto.PictureResponse, error)
	Get(int) (dto.PictureResponse, error)
	GetFile(int) (string, error)
	Create(*multipart.FileHeader) (dto.PictureResponse, error)
	Delete(int) error
	Update(int, *multipart.FileHeader) (dto.PictureResponse, error)
}

type picturesService struct {
	repository repository.PicturesRepository
	storage    storage.Storage
}

func NewPicturesService(repository repository.PicturesRepository, storage storage.Storage) PicturesService {
	return &picturesService{repository, storage}
}

func (s *picturesService) List() ([]dto.PictureResponse, error) {
	pictures, err := s.repository.GetAll()
	var pictureResponses []dto.PictureResponse
	if err != nil {
		return pictureResponses, err
	}

	for _, eachPicture := range pictures {
		pictureResponses = append(pictureResponses, eachPicture.ToPictureResponse())
	}
	return pictureResponses, err
}

func (s *picturesService) Get(id int) (dto.PictureResponse, error) {
	picture, err := s.repository.GetById(id)
	if err != nil {
		return dto.PictureResponse{}, err
	}

	return picture.ToPictureResponse(), nil
}

func (s *picturesService) GetFile(id int) (string, error) {
	picture, err := s.repository.GetById(id)
	if err != nil {
		return "", err
	}

	return picture.Destination, nil
}

func (s *picturesService) Create(file *multipart.FileHeader) (dto.PictureResponse, error) {
	destination, err := s.storage.Save(file)
	if err != nil {
		return dto.PictureResponse{}, err
	}

	var request dto.CreatePictureRequest
	request.Name = file.Filename
	request.Destination = destination

	picture, err := s.repository.Create(request)
	if err != nil {
		return dto.PictureResponse{}, err
	}

	return picture.ToPictureResponse(), nil
}

func (s *picturesService) Delete(id int) error {
	err := s.repository.Delete(id)
	return err
}

func (s *picturesService) Update(id int, file *multipart.FileHeader) (dto.PictureResponse, error) {
	destination, err := s.storage.Save(file)
	if err != nil {
		return dto.PictureResponse{}, err
	}

	var request dto.UpdatePictureRequest
	request.Name = file.Filename
	request.Destination = destination

	picture, err := s.repository.Update(id, request)
	if err != nil {
		return dto.PictureResponse{}, err
	}

	return picture.ToPictureResponse(), nil
}
