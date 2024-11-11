package models

import (
	"time"

	"github.com/google/uuid"
)

type ResearchTitle struct {
	ID          uuid.UUID `db:"id" json:"id"`
	ScientistID uuid.UUID `db:"scientist_id" json:"scientist_id"`
	Title       string    `db:"title" json:"title"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
