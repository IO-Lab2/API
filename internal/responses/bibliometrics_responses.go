package responses

import (
	"time"

	"github.com/google/uuid"
)

type BibliometricsResponse struct {
	Body []BibliometricBody `json:"body" doc:"Bibliometrics object"`
}

type BibliometricBody struct {
	ID               uuid.UUID `json:"id" doc:"Bibliometric ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
	HIndexWos        int       `json:"h_index_wos" doc:"HIndex Wos" format:"int" example:"1"`
	HIndexScopus     int       `json:"h_index_scopus" doc:"HIndex Scopus" format:"int" example:"2"`
	CitationCount    int       `json:"citation_count" doc:"Citation count" format:"int" example:"67"`
	PublicationCount int       `json:"publication_count" doc:"Publication count" format:"int" example:"123"`
	MinisterialScore float64   `json:"ministerial_score" doc:"Ministerial score" format:"float" example:"65.7"`
	ScientistID      uuid.UUID `json:"scientist_id" doc:"Scientist ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
	CreatedAt        time.Time `json:"created_at" doc:"Creation date of bibliometric" format:"date-time" example:"2021-01-01T00:00:00Z"`
	UpdatedAt        time.Time `json:"updated_at" doc:"Update date of bibliometric" format:"date-time" example:"2021-01-01T00:00:00Z"`
}
