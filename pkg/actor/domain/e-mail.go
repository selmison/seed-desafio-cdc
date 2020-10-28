package domain

import (
	"regexp"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Email struct {
	email string
}

// NewEmail creates a new valid email
func NewEmail(email string) (Email, error) {
	if !isEmailValid(email) {
		return Email{}, ErrEmailIsNotValidated
	}
	return Email{email: email}, nil
}

// String implements the fmt.Stringer interface.
func (n Email) String() string {
	return n.email
}

// MarshalText serializes the object.
func (n Email) MarshalText() ([]byte, error) {
	return []byte(n.email), nil
}

// UnmarshalText deserializes the object.
func (n *Email) UnmarshalText(d []byte) error {
	var err error
	*n, err = NewEmail(string(d))
	return err
}

// isEmailValid checks if the email provided passes the required structure and length.
func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}
