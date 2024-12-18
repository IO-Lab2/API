package models

import (
	"time"

	"github.com/google/uuid"
)

type Publication struct {
	ID              uuid.UUID  `db:"id" json:"id"`
	Title           *string    `db:"title" json:"title"`
	Journal         *string    `db:"journal" json:"journal"`
	PublicationDate *time.Time `db:"publication_date" json:"publication_date"`
	ImpactFactor    *float64   `db:"journal_impact_factor" json:"impact_factor"`
	Publisher       *string    `db:"publisher" json:"publisher"`
	JournalType     *string    `db:"journal_type" json:"journal_type"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time  `db:"updated_at" json:"updated_at"`
}
type CreatePublication struct {
	Title           string    `db:"title" json:"title"`
	Journal         string    `db:"journal" json:"journal"`
	PublicationDate time.Time `db:"publication_date" json:"publication_date"`
	ImpactFactor    float64   `db:"impact_factor" json:"impact_factor"`
}
type UpdatePublication struct {
	ID              uuid.UUID `db:"id" json:"id"`
	Title           string    `db:"title" json:"title"`
	Journal         string    `db:"journal" json:"journal"`
	PublicationDate time.Time `db:"publication_date" json:"publication_date"`
	ImpactFactor    float64   `db:"impact_factor" json:"impact_factor"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}
type DeletePublication struct {
	ID uuid.UUID `db:"id" json:"id"`
}
