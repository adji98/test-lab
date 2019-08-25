package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

// Season model struct
type Season struct {
	ID           uuid.UUID    `json:"id" db:"id"`
	CreatedAt    time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at" db:"updated_at"`
	SeriesID     uuid.UUID    `json:"series_id" db:"series_id"`
	SeasonNo     int          `json:"season_no" db:"season_no"`
	Plot         nulls.String `json:"plot" db:"plot"`
	Completed    bool         `json:"completed" db:"completed"`
	EpisodeCount int          `json:"episode_count" db:"episode_count"`
}

// String is not required by pop and may be deleted
func (s Season) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Seasons is not required by pop and may be deleted
type Seasons []Season

// String is not required by pop and may be deleted
func (s Seasons) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Season) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: s.SeasonNo, Name: "SeasonNo"},
		&validators.IntIsPresent{Field: s.EpisodeCount, Name: "EpisodeCount"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Season) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Season) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
