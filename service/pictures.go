package service

import "github.com/VinayakBagaria/go-cat-pictures/repository"

type PicturesService interface {
}

type picturesService struct {
	picturesRepository repository.PicturesRepository
}

func NewPicturesService(picturesRepository repository.PicturesRepository) PicturesService {
	return &picturesService{picturesRepository}
}
