package models

import (
	"time"

	"github.com/google/uuid"
)

type Bibliometrics struct {
	ID               uuid.UUID `db:"id" json:"id"`
	HIndex           int       `db:"h_index" json:"h_index"`
	CitationCount    int       `db:"citation_count" json:"citation_count"`
	PublicationCount int       `db:"publication_count" json:"publication_count"`
	MinisterialScore float64   `db:"ministerial_score" json:"ministerial_score"`
	ScientistID      uuid.UUID `db:"scientist_id" json:"scientist_id"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
}
