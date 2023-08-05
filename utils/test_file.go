package utils

import "mime/multipart"

func NewTestFile(fileName string) *multipart.FileHeader {
	return &multipart.FileHeader{
		Filename: fileName,
		Size:     1000,
	}
}
