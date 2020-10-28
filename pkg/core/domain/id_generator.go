package domain

import "github.com/google/uuid"

// IdGenerator generates a new ID.
type IdGenerator interface {
	// Generate generates a new ID.
	Generate() (string, error)
}

type generator struct{}

func NewGenerator() IdGenerator {
	return generator{}
}

func (g generator) Generate() (string, error) {
	return uuid.New().String(), nil
}
