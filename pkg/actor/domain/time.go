package domain

import (
	"encoding/json"
	"time"
)

type Time struct {
	time time.Time
}

// GenerateTime generates a new valid time.
func GenerateTime() Time {
	return Time{time: time.Now()}
}

func (t *Time) Validate() error {
	if t.time.IsZero() {
		return ErrTimeIsNotValidated
	}
	return nil
}

// MarshalJSON serializes the object.
func (t Time) MarshalJSON() ([]byte, error) {
	if err := t.Validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t.time)
}

// UnmarshalJSON deserializes the object.
func (t *Time) UnmarshalJSON(d []byte) error {
	return json.Unmarshal(d, &t.time)
}
