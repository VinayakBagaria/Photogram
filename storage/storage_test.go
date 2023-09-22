package storage

import (
	"os"
	"testing"

	"github.com/VinayakBagaria/photogram/utils"
	"github.com/stretchr/testify/assert"
)

func TestStorageCreation(t *testing.T) {
	path := "./test_images_storage"
	os.RemoveAll(path)

	storage := NewStorage(path)
	fileName := utils.NewUniqueString() + ".txt"

	assert.Equal(t, storage.GetFullPath(fileName), path+"/"+fileName)
	os.RemoveAll(path)
}

func TestStorageRetrieval(t *testing.T) {
	storage := NewStorage("./")
	data, err := storage.Get("storage_test.go")
	assert.Nil(t, err)
	assert.Greater(t, len(data), 0)
}
