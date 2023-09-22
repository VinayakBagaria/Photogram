package service

import (
	"errors"
	"time"

	"github.com/VinayakBagaria/photogram/db"
	"github.com/VinayakBagaria/photogram/dto"
)

type fakeRepository struct {
	data map[int]*db.Picture
}

func NewFakeRepository() *fakeRepository {
	return &fakeRepository{
		data: map[int]*db.Picture{},
	}
}

func (f *fakeRepository) Create(request *dto.PictureRequest) (*db.Picture, error) {
	rowId := len(f.data) + 1
	picture := &db.Picture{
		ID:          uint(rowId),
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
	f.data[rowId] = picture
	return picture, nil
}

func (f *fakeRepository) Update(id int, request *dto.PictureRequest) (*db.Picture, error) {
	rowId := uint(id)
	for _, eachRow := range f.data {
		if eachRow.ID == rowId {
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
			f.data[id] = updatedPicture
			return updatedPicture, nil
		}
	}

	return nil, errors.New("unable to find")
}

func (f *fakeRepository) Delete(id int) error {
	if _, ok := f.data[id]; ok {
		delete(f.data, id)
		return nil
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

	keys := []int{}
	for eachKey := range f.data {
		keys = append(keys, eachKey)
	}

	limitedKeys := keys[start:end]
	response := []*db.Picture{}
	for _, eachKey := range limitedKeys {
		response = append(response, f.data[eachKey])
	}

	return response, int64(len(f.data)), nil
}

func (f *fakeRepository) GetById(id int) (*db.Picture, error) {
	if val, ok := f.data[id]; ok {
		return val, nil
	}
	return nil, errors.New("unable to find")
}
