package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

// Series model struct
type Series struct {
	ID           uuid.UUID `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	Title        string    `json:"title" db:"title"`
	Plot         string    `json:"plot" db:"plot"`
	Completed    bool      `json:"completed" db:"completed"`
	SeasonCount  int       `json:"season_count" db:"season_count"`
	EpisodeCount int       `json:"episode_count" db:"episode_count"`
}

// String is not required by pop and may be deleted
func (s Series) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Series) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: s.Title, Name: "Title"},
		&validators.StringIsPresent{Field: s.Plot, Name: "Plot"},
		&validators.IntIsPresent{Field: s.SeasonCount, Name: "SeasonCount"},
		&validators.IntIsPresent{Field: s.EpisodeCount, Name: "EpisodeCount"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Series) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Series) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
