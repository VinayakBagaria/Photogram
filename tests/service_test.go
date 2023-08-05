package tests

import (
	"mime/multipart"
	"reflect"
	"strings"
	"testing"

	"github.com/VinayakBagaria/go-cat-pictures/dto"
	"github.com/VinayakBagaria/go-cat-pictures/service"
)

func newFile(fileName string) *multipart.FileHeader {
	return &multipart.FileHeader{
		Filename: fileName,
		Size:     1000,
	}
}

func TestServiceFunctions(t *testing.T) {
	repo := NewFakeRepository()
	storage := NewFakeStorage()
	svc := service.NewPicturesService(repo, storage)

	t.Run("create entry", func(t *testing.T) {
		file := newFile(NewUniqueString())
		createResponse, errorState := svc.Create(file)
		if errorState != nil {
			assertNotNull(t, errorState.Error)
		}

		assertData(t, strings.HasSuffix(createResponse.Name, file.Filename), true)
		assertData(t, createResponse.Id, repo.data[0].ID)

		entryId := int(createResponse.Id)
		fileResponse, err := svc.Get(entryId)
		assertNull(t, err)
		assertData(t, fileResponse.Name, createResponse.Name)
	})

	t.Run("update entry", func(t *testing.T) {
		file := newFile(NewUniqueString())
		updateResponse, errorState := svc.Update(int(repo.data[0].ID), file)

		if errorState != nil {
			assertNotNull(t, errorState.Error)
		}

		assertData(t, true, strings.HasSuffix(updateResponse.Name, file.Filename))
		fileResponse, err := svc.Get(int(updateResponse.Id))
		assertNull(t, err)
		assertData(t, fileResponse.Name, updateResponse.Name)
	})

	t.Run("list page", func(t *testing.T) {
		listResponse, count, err := svc.List(10, 1)
		totalCount := int(count)

		assertNull(t, err)
		assertData(t, totalCount, len(listResponse))
		assertData(t, totalCount, len(repo.data))
		for index, eachResponse := range listResponse {
			assertData(t, eachResponse, repo.data[index].ToPictureResponse())
		}
	})

	t.Run("out of bounds list page", func(t *testing.T) {
		invalidPage := len(repo.data) + 1
		listResponse, count, err := svc.List(1, invalidPage)
		totalCount := int(count)

		assertNull(t, err)
		assertData(t, totalCount, len(repo.data))
		assertData(t, listResponse, []*dto.PictureResponse{})
	})

	t.Run("get entry", func(t *testing.T) {
		randomEntry := NewRandomNumber(1, len(repo.data))
		response, err := svc.Get(randomEntry)

		assertNull(t, err)
		assertData(t, response, repo.data[randomEntry-1].ToPictureResponse())
	})

	t.Run("invalid get entry", func(t *testing.T) {
		_, err := svc.Get(-1)

		assertNotNull(t, err)
	})

	t.Run("delete entry", func(t *testing.T) {
		initialLength := len(repo.data)
		randomEntry := NewRandomNumber(1, initialLength)
		err := svc.Delete(randomEntry)

		assertNull(t, err)
		assertData(t, len(repo.data), initialLength-1)
	})

	t.Run("invalid delete entry", func(t *testing.T) {
		err := svc.Delete(-1)

		assertNotNull(t, err)
	})

}

func assertNotNull(t *testing.T, val any) {
	t.Helper()
	if val == nil {
		t.Fatal(val)
	}
}

func assertNull(t *testing.T, val any) {
	t.Helper()
	if val != nil {
		t.Fatal(val)
	}
}

func assertData(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
