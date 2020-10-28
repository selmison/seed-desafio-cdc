package domain

// UpdateActorDTO contains updates of an existing actor.
type UpdateActorDTO struct {
	Name        *Name  `json:"name"`
	Email       *Email `json:"e-mail"`
	Description *Desc  `json:"description"`
}

func (a UpdateActorDTO) IsEmpty() bool {
	return a.Compare(UpdateActorDTO{})
}

func (a UpdateActorDTO) Compare(b UpdateActorDTO) bool {
	if &a == &b {
		return true
	}
	if a.Name != b.Name {
		return false
	}
	if a.Email != b.Email {
		return false
	}
	if a.Description != b.Description {
		return false
	}
	return true
}
