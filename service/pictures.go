package service

import (
	"mime/multipart"
	"net/http"

	"github.com/VinayakBagaria/go-cat-pictures/dto"
	"github.com/VinayakBagaria/go-cat-pictures/repository"
	"github.com/VinayakBagaria/go-cat-pictures/storage"
)

type PicturesService interface {
	Create(*multipart.FileHeader) (*dto.PictureResponse, *dto.InvalidPictureFileError)
	Update(int, *multipart.FileHeader) (*dto.PictureResponse, *dto.InvalidPictureFileError)
	List() ([]*dto.PictureResponse, error)
	Get(int) (*dto.PictureResponse, error)
	GetFile(int) (string, error)
	Delete(int) error
}

type picturesService struct {
	repository repository.PicturesRepository
	storage    storage.ImageStorage
}

func NewPicturesService(repository repository.PicturesRepository, storage storage.ImageStorage) PicturesService {
	return &picturesService{repository, storage}
}

func (s *picturesService) Create(file *multipart.FileHeader) (*dto.PictureResponse, *dto.InvalidPictureFileError) {
	requestData, createError := s.storage.Save(file)
	if createError != nil {
		return nil, createError
	}

	picture, err := s.repository.Create(requestData)
	if err != nil {
		return nil, &dto.InvalidPictureFileError{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
	}

	return picture.ToPictureResponse(), nil
}

func (s *picturesService) Update(id int, file *multipart.FileHeader) (*dto.PictureResponse, *dto.InvalidPictureFileError) {
	requestData, createError := s.storage.Save(file)
	if createError != nil {
		return nil, createError
	}

	picture, err := s.repository.Update(id, requestData)
	if createError != nil {
		return nil, &dto.InvalidPictureFileError{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
	}

	return picture.ToPictureResponse(), nil
}

func (s *picturesService) List() ([]*dto.PictureResponse, error) {
	pictures, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	pictureResponses := make([]*dto.PictureResponse, 0, len(pictures))
	for _, eachPicture := range pictures {
		pictureResponses = append(pictureResponses, eachPicture.ToPictureResponse())
	}
	return pictureResponses, err
}

func (s *picturesService) Get(id int) (*dto.PictureResponse, error) {
	picture, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return picture.ToPictureResponse(), nil
}

func (s *picturesService) GetFile(id int) (string, error) {
	picture, err := s.repository.GetById(id)
	if err != nil {
		return "", err
	}

	return s.storage.GetFullPath(picture.Destination), nil
}

func (s *picturesService) Delete(id int) error {
	err := s.repository.Delete(id)
	return err
}
