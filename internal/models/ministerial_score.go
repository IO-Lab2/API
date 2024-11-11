package models

import (
	"time"

	"github.com/google/uuid"
)

type MinisterialScore struct {
	ID          uuid.UUID `db:"id" json:"id"`
	ScientistID uuid.UUID `db:"scientist_id" json:"scientist_id"`
	Score       int       `db:"score" json:"score"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
