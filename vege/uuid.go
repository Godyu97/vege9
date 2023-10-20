package vege

import (
	"github.com/google/uuid"
)

//go get github.com/google/uuid

// google uuid v4
func NewGoogleUUID() string {
	return uuid.NewString()
}
