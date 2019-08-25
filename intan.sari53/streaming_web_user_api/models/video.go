package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

// Video model struct
type Video struct {
	ID           uuid.UUID `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	Title        string    `json:"title" db:"title"`
	Plot         string    `json:"plot" db:"plot"`
	Synopsis     string    `json:"synopsis" db:"synopsis"`
	Length       int       `json:"length" db:"length"`
	ViewCount    int       `json:"view_count" db:"view_count"`
	LikeCount    int       `json:"like_count" db:"like_count"`
	DislikeCount int       `json:"dislike_count" db:"dislike_count"`
}

// String is not required by pop and may be deleted
func (v Video) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}

// Videoes is not required by pop and may be deleted
type Videoes []Video

// String is not required by pop and may be deleted
func (v Videoes) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (v *Video) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: v.Title, Name: "Title"},
		&validators.StringIsPresent{Field: v.Plot, Name: "Plot"},
		&validators.StringIsPresent{Field: v.Synopsis, Name: "Synopsis"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (v *Video) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (v *Video) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
