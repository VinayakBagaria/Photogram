package tests

import "github.com/google/uuid"

func NewUniqueString() string {
	return uuid.New().String()
}
