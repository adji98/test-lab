package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
	"time"
)

// VideoView model struct
type VideoView struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	VideoID   uuid.UUID `json:"video_id" db:"video_id"`
	QualityID uuid.UUID `json:"quality_id" db:"quality_id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	Completed bool      `json:"completed" db:"completed"`
	PausedOn  int       `json:"paused_on" db:"paused_on"`
}

// String is not required by pop and may be deleted
func (v VideoView) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}

// VideoViews is not required by pop and may be deleted
type VideoViews []VideoView

// String is not required by pop and may be deleted
func (v VideoViews) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (v *VideoView) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: v.PausedOn, Name: "PausedOn"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (v *VideoView) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (v *VideoView) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
