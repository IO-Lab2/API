package models

import (
	"time"

	"github.com/google/uuid"
)

type CitationCount struct {
	ID          uuid.UUID `db:"id" json:"id"`
	ScientistID uuid.UUID `db:"scientist_id" json:"scientist_id"`
	Count       int       `db:"count" json:"count"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
