package utils

import (
	"math/rand"

	"github.com/google/uuid"
)

func NewUniqueString() string {
	return uuid.New().String()
}

func NewRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}
