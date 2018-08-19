package models

import validation "github.com/go-ozzo/ozzo-validation"

// Supplement represents a supplement record
type Supplement struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// Validate validates the Artist fields.
func (m Supplement) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 120)),
	)
}
