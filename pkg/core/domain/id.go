package domain

import "strings"

type Id struct {
	id string
}

// GenerateId generates a new valid id
func GenerateId() (Id, error) {
	id, err := NewGenerator().Generate()
	if err != nil {
		return Id{}, err
	}
	return NewId(id)
}

// NewId creates a new valid id
func NewId(id string) (Id, error) {
	if strings.TrimSpace(id) == "" {
		return Id{}, ErrIdCouldNotBeEmpty
	}
	return Id{id: id}, nil
}

// String implements the fmt.Stringer interface.
func (i Id) String() string {
	return i.id
}

// MarshalText serializes the object.
func (i Id) MarshalText() ([]byte, error) {
	return []byte(i.id), nil
}

// UnmarshalText deserializes the object.
func (i *Id) UnmarshalText(d []byte) error {
	var err error
	*i, err = NewId(string(d))
	return err
}
