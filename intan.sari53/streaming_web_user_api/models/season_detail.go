package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
	"time"
)

// SeasonDetail model struct
type SeasonDetail struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	SeasonID  uuid.UUID `json:"season_id" db:"season_id"`
	VideoID   uuid.UUID `json:"video_id" db:"video_id"`
	EpisodeNo int       `json:"episode_no" db:"episode_no"`
}

// String is not required by pop and may be deleted
func (s SeasonDetail) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// SeasonDetails is not required by pop and may be deleted
type SeasonDetails []SeasonDetail

// String is not required by pop and may be deleted
func (s SeasonDetails) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *SeasonDetail) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: s.EpisodeNo, Name: "EpisodeNo"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *SeasonDetail) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *SeasonDetail) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
