package tests

import (
	"errors"
	"time"

	"github.com/VinayakBagaria/go-cat-pictures/db"
	"github.com/VinayakBagaria/go-cat-pictures/dto"
)

type fakeRepository struct {
	data []*db.Picture
}

func NewFakeRepository() *fakeRepository {
	return &fakeRepository{
		data: []*db.Picture{},
	}
}

func (f *fakeRepository) Create(request *dto.PictureRequest) (*db.Picture, error) {
	picture := &db.Picture{
		ID:          uint(len(f.data) + 1),
		CreatedOn:   time.Now().Unix(),
		UpdatedOn:   time.Now().Unix(),
		Deleted:     false,
		Name:        request.Name,
		Destination: request.Destination,
		Height:      request.Height,
		Width:       request.Width,
		Size:        request.Size,
		ContentType: request.ContentType,
	}
	f.data = append(f.data, picture)
	return picture, nil
}

func (f *fakeRepository) Update(id int, request *dto.PictureRequest) (*db.Picture, error) {
	for index, eachRow := range f.data {
		if eachRow.ID == uint(id) {
			updatedPicture := &db.Picture{
				ID:        eachRow.ID,
				CreatedOn: eachRow.CreatedOn,
				UpdatedOn: time.Now().Unix(),
				Deleted:   false,

				Name:        request.Name,
				Destination: request.Destination,
				Height:      request.Height,
				Width:       request.Width,
				Size:        request.Size,
				ContentType: request.ContentType,
			}
			f.data[index] = updatedPicture
			return updatedPicture, nil
		}
	}

	return nil, errors.New("unable to find")
}

func (f *fakeRepository) Delete(id int) error {
	for index, eachRow := range f.data {
		if eachRow.ID == uint(id) {
			f.data = append(f.data[:index], f.data[index+1:]...)
			return nil
		}
	}
	return errors.New("unable to find")
}

func (f *fakeRepository) GetAll(limit, page int) ([]*db.Picture, int64, error) {
	start := (page - 1) * limit
	end := start + limit + 1

	if start >= len(f.data) {
		return []*db.Picture{}, int64(len(f.data)), nil
	}

	if end > len(f.data) {
		end = len(f.data)
	}

	return f.data[start:end], int64(len(f.data)), nil
}

func (f *fakeRepository) GetById(id int) (*db.Picture, error) {
	for _, eachRow := range f.data {
		if eachRow.ID == uint(id) {
			return eachRow, nil
		}
	}
	return nil, errors.New("unable to find")
}
