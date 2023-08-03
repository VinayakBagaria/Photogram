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
	Save(*multipart.FileHeader) (string, error)
	Get(string) ([]byte, error)
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
	path := s.path + "/" + destination

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	out, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return destination, err
}

func (s *localStorage) Get(destination string) ([]byte, error) {
	path := s.path + "/" + destination
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body, err := ioutil.ReadAll(file)
	return body, err
}
