package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type Entrie struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	Key       string    `json:"key" db:"key"`
	Value     string    `json:"value" db:"value"`
	Category  string    `json:"category" db:"category"`
}

// String is not required by pop and may be deleted
func (e Entrie) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Entries is not required by pop and may be deleted
type Entries []Entrie

// String is not required by pop and may be deleted
func (e Entries) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Validate gets run everytime you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (e *Entrie) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: e.Key, Name: "Key"},
		&validators.StringIsPresent{Field: e.Value, Name: "Value"},
		&validators.StringIsPresent{Field: e.Category, Name: "Category"},
	), nil
}

// ValidateSave gets run everytime you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (e *Entrie) ValidateSave(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run everytime you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (e *Entrie) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
