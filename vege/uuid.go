package vege

import (
	"github.com/google/uuid"
)

//go get github.com/google/uuid

func NewGoogleUUID() string {
	return uuid.NewString()
}
