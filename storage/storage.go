package storage

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type Storage interface {
	Save(*multipart.FileHeader) (string, error)
}

type localStorage struct {
	path string
}

func NewStorage(path string) Storage {
	return &localStorage{path}
}

func (s *localStorage) Save(file *multipart.FileHeader) (string, error) {
	extension := filepath.Ext(file.Filename)
	destination := uuid.New().String() + extension

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	out, err := os.Create(s.path + "/" + destination)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return destination, err
}
