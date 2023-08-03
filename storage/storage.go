package storage

import (
	"errors"
	"image/png"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/VinayakBagaria/go-cat-pictures/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var ALLOWED_CONTENT_TYPES = [6]string{"image/jpeg", "image/png", "image/gif", "image/tiff", "image/bmp", "video/webm"}

type ImageStorage interface {
	GetFullPath(string) string
	Save(*multipart.FileHeader) (*dto.PictureRequest, *dto.InvalidPictureFileError)
	Get(string) ([]byte, error)
}

type localImageStorage struct {
	path string
}

func NewStorage(path string) ImageStorage {
	return &localImageStorage{path}
}

func (s *localImageStorage) GetFullPath(destination string) string {
	return s.path + "/" + destination
}

func (s *localImageStorage) Save(file *multipart.FileHeader) (*dto.PictureRequest, *dto.InvalidPictureFileError) {
	extension := filepath.Ext(file.Filename)
	destination := uuid.New().String() + extension
	fullPath := s.GetFullPath(destination)

	src, err := file.Open()
	if err != nil {
		return nil, &dto.InvalidPictureFileError{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
	}
	defer src.Close()

	buffer := make([]byte, 512)
	_, err = src.Read(buffer)
	if err != nil {
		return nil, &dto.InvalidPictureFileError{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
	}

	fileType := http.DetectContentType(buffer)
	isImage := false
	for _, desiredContentType := range ALLOWED_CONTENT_TYPES {
		if fileType == desiredContentType {
			isImage = true
			break
		}
	}

	if !isImage {
		return nil, &dto.InvalidPictureFileError{
			StatusCode: http.StatusBadRequest,
			Error:      errors.New("unsupported format"),
			Data:       gin.H{"format": fileType},
		}
	}

	_, err = src.Seek(0, io.SeekStart)
	if err != nil {
		return nil, &dto.InvalidPictureFileError{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
	}

	im, err := png.DecodeConfig(src)
	if err != nil {
		return nil, &dto.InvalidPictureFileError{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
			Data:       gin.H{"format": fileType},
		}
	}

	out, err := os.Create(fullPath)
	if err != nil {
		return nil, &dto.InvalidPictureFileError{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
	}
	defer out.Close()

	src.Seek(0, io.SeekStart)
	_, err = io.Copy(out, src)
	if err != nil {
		return nil, &dto.InvalidPictureFileError{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
	}

	pictureFile := &dto.PictureRequest{
		Name:        file.Filename,
		Destination: destination,
		Height:      int32(im.Height),
		Width:       int32(im.Width),
		Size:        int32(file.Size),
		ContentType: fileType,
	}

	return pictureFile, nil
}

func (s *localImageStorage) Get(destination string) ([]byte, error) {
	fullPath := s.GetFullPath(destination)
	file, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body, err := ioutil.ReadAll(file)
	return body, err
}
