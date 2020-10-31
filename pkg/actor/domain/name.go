package domain

import (
	"strings"
)

type Name struct {
	name string
}

// NewName creates a new valid name
func NewName(name string) (Name, error) {
	n := Name{name: name}
	if err := n.Validate(); err != nil {
		return Name{}, err
	}
	return n, nil
}

func (n *Name) Validate() error {
	if strings.TrimSpace(n.name) == "" {
		return ErrNameCouldNotBeEmpty
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (n Name) String() string {
	return n.name
}

// MarshalText serializes the object.
func (n Name) MarshalText() ([]byte, error) {
	if err := n.Validate(); err != nil {
		return nil, err
	}
	return []byte(n.name), nil
}

// UnmarshalText deserializes the object.
func (n *Name) UnmarshalText(d []byte) error {
	var err error
	*n, err = NewName(string(d))
	return err
}
