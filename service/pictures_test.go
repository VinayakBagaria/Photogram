package service

import (
	"reflect"
	"strings"
	"testing"

	"github.com/VinayakBagaria/go-cat-pictures/dto"
	"github.com/VinayakBagaria/go-cat-pictures/utils"
	"github.com/stretchr/testify/assert"
)

func TestServiceFunctions(t *testing.T) {
	repo := NewFakeRepository()
	storage := NewFakeStorage()
	svc := NewPicturesService(repo, storage)

	t.Run("create entry", func(t *testing.T) {
		file := utils.NewTestFile(utils.NewUniqueString())
		createResponse, errorState := svc.Create(file)
		if errorState != nil {
			assert.NotNil(t, errorState.Error)
		}

		assert.Equal(t, strings.HasSuffix(createResponse.Name, file.Filename), true)
		value, existsInStorage := repo.data[int(createResponse.Id)]
		assert.Equal(t, existsInStorage, true)
		assert.Equal(t, createResponse, value.ToPictureResponse())

		entryId := int(createResponse.Id)
		fileResponse, err := svc.Get(entryId)
		assert.Nil(t, err)
		assert.Equal(t, fileResponse.Name, createResponse.Name)
	})

	t.Run("update entry", func(t *testing.T) {
		file := utils.NewTestFile(utils.NewUniqueString())

		allKeys := reflect.ValueOf(repo.data).MapKeys()
		randomKey := int(allKeys[utils.NewRandomNumber(0, len(allKeys)-1)].Int())

		updateResponse, errorState := svc.Update(int(repo.data[randomKey].ID), file)

		if errorState != nil {
			assert.NotNil(t, errorState.Error)
		}

		assert.Equal(t, true, strings.HasSuffix(updateResponse.Name, file.Filename))
		fileResponse, err := svc.Get(int(updateResponse.Id))
		assert.Nil(t, err)
		assert.Equal(t, fileResponse.Name, updateResponse.Name)
	})

	t.Run("list page", func(t *testing.T) {
		listResponse, count, err := svc.List(10, 1)
		totalCount := int(count)

		assert.Nil(t, err)
		assert.Equal(t, totalCount, len(listResponse))
		assert.Equal(t, totalCount, len(repo.data))
		for _, eachResponse := range listResponse {
			assert.Equal(t, eachResponse, repo.data[int(eachResponse.Id)].ToPictureResponse())
		}
	})

	t.Run("out of bounds list page", func(t *testing.T) {
		invalidPage := len(repo.data) + 1
		listResponse, count, err := svc.List(1, invalidPage)
		totalCount := int(count)

		assert.Nil(t, err)
		assert.Equal(t, totalCount, len(repo.data))
		assert.Equal(t, listResponse, []*dto.PictureResponse{})
	})

	t.Run("get entry", func(t *testing.T) {
		randomEntry := utils.NewRandomNumber(1, len(repo.data))
		response, err := svc.Get(randomEntry)

		assert.Nil(t, err)
		assert.Equal(t, response, repo.data[int(response.Id)].ToPictureResponse())
	})

	t.Run("invalid get entry", func(t *testing.T) {
		_, err := svc.Get(-1)

		assert.NotNil(t, err)
	})

	t.Run("delete entry", func(t *testing.T) {
		initialLength := len(repo.data)
		randomEntry := utils.NewRandomNumber(1, initialLength)
		err := svc.Delete(randomEntry)

		assert.Nil(t, err)
		assert.Equal(t, len(repo.data), initialLength-1)
	})

	t.Run("invalid delete entry", func(t *testing.T) {
		err := svc.Delete(-1)

		assert.NotNil(t, err)
	})

}
