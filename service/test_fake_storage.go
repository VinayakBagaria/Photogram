package service

import (
	"errors"
	"mime/multipart"
	"path/filepath"

	"github.com/VinayakBagaria/go-cat-pictures/dto"
	"github.com/VinayakBagaria/go-cat-pictures/storage"
	"github.com/VinayakBagaria/go-cat-pictures/utils"
)

type fakeStorage struct {
	BaseDirectory string
	Contents      map[string][]byte
}

func NewFakeStorage() storage.ImageStorage {
	return &fakeStorage{
		BaseDirectory: "/some-place",
		Contents:      make(map[string][]byte),
	}
}

func (s *fakeStorage) GetFullPath(destination string) string {
	return s.BaseDirectory + "/" + destination
}

func (s *fakeStorage) Save(file *multipart.FileHeader) (*dto.PictureRequest, *dto.InvalidPictureFileError) {
	randomFileName := utils.NewUniqueString() + "----" + file.Filename
	destination := randomFileName + filepath.Ext(file.Filename)
	pictureFile := &dto.PictureRequest{
		Name:        randomFileName,
		Destination: s.GetFullPath(destination),
		Height:      100,
		Width:       100,
		Size:        int32(file.Size),
		ContentType: "image/jpeg",
	}
	s.Contents[destination] = []byte(pictureFile.Name)
	return pictureFile, nil
}

func (s *fakeStorage) Get(destination string) ([]byte, error) {
	if val, ok := s.Contents[destination]; ok {
		return val, nil
	}
	return nil, errors.New("unable to find")
}
