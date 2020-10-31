package domain

import "strings"

type Desc struct {
	description string
}

// NewDesc creates a new valid description
func NewDesc(description string) (Desc, error) {
	description = strings.TrimSpace(description)
	if len(description) == 0 {
		return Desc{}, ErrDescIsNotValidated
	}
	return Desc{description: description}, nil
}

// String implements the fmt.Stringer interface.
func (n Desc) String() string {
	return n.description
}

// MarshalText serializes the object.
func (n Desc) MarshalText() ([]byte, error) {
	return []byte(n.description), nil
}

// UnmarshalText deserializes the object.
func (n *Desc) UnmarshalText(d []byte) error {
	var err error
	*n, err = NewDesc(string(d))
	return err
}
