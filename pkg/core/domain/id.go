package domain

import (
	"strings"

	"github.com/google/uuid"
)

type ID struct {
	id string
}

// GenerateID generates a new valid id
func GenerateID() (ID, error) {
	id := uuid.New().String()
	return NewID(id)
}

// NewID creates a new valid id
func NewID(id string) (ID, error) {
	if strings.TrimSpace(id) == "" {
		return ID{}, ErrIdCouldNotBeEmpty
	}
	return ID{id: id}, nil
}

// String implements the fmt.Stringer interface.
func (i ID) String() string {
	return i.id
}

// MarshalText serializes the object.
func (i ID) MarshalText() ([]byte, error) {
	return []byte(i.id), nil
}

// UnmarshalText deserializes the object.
func (i *ID) UnmarshalText(d []byte) error {
	var err error
	*i, err = NewID(string(d))
	return err
}
