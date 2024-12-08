package responses

import (
	"time"

	"github.com/google/uuid"
)

type PublicationResponse struct {
	Body *PublicationBody `json:"publication_body"`
}

type PublicationsResponse struct {
	Body []PublicationBody
}
type PublicationBody struct {
	ID                  uuid.UUID `db:"id" json:"id" doc:"Publication ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
	Title               string    `db:"title" json:"title" doc:"Publication title" format:"string" example:"Quantum Mechanics and Path Integrals"`
	Journal             *string   `db:"journal" json:"journal" doc:"Journal where publication is published" format:"string" example:"Nature"`
	PublicationDate     time.Time `db:"publication_date" json:"publication_date" doc:"Publication date" format:"date-time" example:"2021-01-01T00:00:00Z"`
	JournalImpactFactor float64   `db:"journal_impact_factor" json:"journal_impact_factor" doc:"Impact factor" format:"float" example:"12.123"`
	CreatedAt           time.Time `db:"created_at" json:"created_at" doc:"Creation date of publication" format:"date-time" example:"2021-01-01T00:00:00Z"`
	UpdatedAt           time.Time `db:"updated_at" json:"updated_at" doc:"Update date of publication" format:"date-time" examle:"2021-01-01T00:00:00Z"`
}
type CreatePublication struct {
	ID uuid.UUID `db:"id" json:"id" doc:"ID of created publication" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
}
type CreatePublicationResponse struct {
	Body CreatePublication `json:"body" doc:"Publication creation object"`
}
