package responses

import (
	"time"

	"github.com/google/uuid"
)

type PublicationsResponse struct {
	Body []PublicationBody
}
type PublicationBody struct {
	ID              uuid.UUID `json:"id" doc:"Publication ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
	Title           string    `json:"title" doc:"Publication title" format:"string" example:"Quantum Mechanics and Path Integrals"`
	Journal         string    `json:"journal" doc:"Journal where publication is published" format:"string" example:"Nature"`
	PublicationDate time.Time `json:"publication_date" doc:"Publication date" format:"date-time" example:"2021-01-01T00:00:00Z"`
	CitationsCount  int       `json:"citations_count" doc:"Citations count" format:"int" example:"12"`
	ImpactFactor    float64   `json:"impact_factor" doc:"Impact factor" format:"float" example:"12.123"`
	CreatedAt       time.Time `json:"created_at" doc:"Creation date of publication" format:"date-time" example:"2021-01-01T00:00:00Z"`
	UpdatedAt       time.Time `json:"updated_at" doc:"Update date of publication" format:"date-time" examle:"2021-01-01T00:00:00Z"`
}
