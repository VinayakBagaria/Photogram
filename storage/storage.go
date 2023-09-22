package storage

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/VinayakBagaria/photogram/dto"
	"github.com/VinayakBagaria/photogram/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"golang.org/x/image/webp"
)

var CONTENT_DECODERS = map[string](func(r io.Reader) (image.Config, error)){
	"image/jpeg": jpeg.DecodeConfig,
	"image/png":  png.DecodeConfig,
	"image/gif":  gif.DecodeConfig,
	"image/tiff": tiff.DecodeConfig,
	"image/webp": webp.DecodeConfig,
	"image/bmp":  bmp.DecodeConfig,
}

type ImageStorage interface {
	GetFullPath(string) string
	Save(*multipart.FileHeader) (*dto.PictureRequest, *dto.InvalidPictureFileError)
	Get(string) ([]byte, error)
}

type localImageStorage struct {
	path string
}

func NewStorage(path string) ImageStorage {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Fatalln("Unable to make directory: %w", path)
		}
	}

	return &localImageStorage{path}
}

func (s *localImageStorage) GetFullPath(destination string) string {
	return s.path + "/" + destination
}

func (s *localImageStorage) Save(file *multipart.FileHeader) (*dto.PictureRequest, *dto.InvalidPictureFileError) {
	extension := filepath.Ext(file.Filename)
	destination := utils.NewUniqueString() + extension
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
	decoder, ok := CONTENT_DECODERS[fileType]
	if !ok {
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

	imageConfig, err := decoder(src)
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
		Height:      int32(imageConfig.Height),
		Width:       int32(imageConfig.Width),
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
