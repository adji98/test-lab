package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type ContinueView struct {
	UserID      uuid.UUID `json:"user_id" db:"user_id" rw:"r"`
	VideoID     uuid.UUID `json:"video_id" db:"video_id" rw:"r"`
	QualityID   uuid.UUID `json:"quality_id" db:"quality_id" rw:"r"`
	PausedOn    int       `json:"paused_on" db:"paused_on" rw:"r"`
	LastView    time.Time `json:"last_view_at" db:"last_view_at" rw:"r"`
	Title       string    `json:"title" db:"title" rw:"r"`
	SeasonNo    int       `json:"season_no" db:"season_no" rw:"r"`
	SeriesTitle string    `json:"series_title" db:"series_title" rw:"r"`
}

type ContinueViews []ContinueView
