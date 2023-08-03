package storage

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type Storage interface {
	GetFullPath(string) string
	Save(*multipart.FileHeader) (string, error)
	Get(string) ([]byte, error)
}

type localStorage struct {
	path string
}

func NewStorage(path string) Storage {
	return &localStorage{path}
}

func (s *localStorage) GetFullPath(destination string) string {
	return s.path + "/" + destination
}

func (s *localStorage) Save(file *multipart.FileHeader) (string, error) {
	extension := filepath.Ext(file.Filename)
	destination := uuid.New().String() + extension
	fullPath := s.GetFullPath(destination)

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	out, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return destination, err
}

func (s *localStorage) Get(destination string) ([]byte, error) {
	fullPath := s.GetFullPath(destination)
	file, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body, err := ioutil.ReadAll(file)
	return body, err
}
