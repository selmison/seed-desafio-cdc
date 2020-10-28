package domain

// NewActorDTO contains the details of a new Actor.
type NewActorDTO struct {
	Name        Name  `json:"name"`
	Email       Email `json:"e-mail"`
	Description Desc  `json:"description"`
}
