package responses

import (
	"time"

	"github.com/google/uuid"
)

type ListOfBibliometricsResponse struct {
	Body []BibliometricBody `json:"body" doc:"Bibliometrics object"`
}

type BibliometricsResponse struct {
	Body *BibliometricBody `json:"body" doc:"Bibliometrics object"`
}

type BibliometricBody struct {
	ID               uuid.UUID `db:"id" json:"id" doc:"Bibliometric ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
	HIndexWos        int       `db:"h_index_wos" json:"h_index_wos" doc:"HIndex Wos" format:"int" example:"1"`
	HIndexScopus     int       `db:"h_index_scopus" json:"h_index_scopus" doc:"HIndex Scopus" format:"int" example:"2"`
	PublicationCount int       `db:"publication_count" json:"publication_count" doc:"Publication count" format:"int" example:"123"`
	MinisterialScore float64   `db:"ministerial_score" json:"ministerial_score" doc:"Ministerial score" format:"float" example:"65.7"`
	ScientistID      uuid.UUID `db:"scientist_id" json:"scientist_id" doc:"Scientist ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
	CreatedAt        time.Time `db:"created_at" json:"created_at" doc:"Creation date of bibliometric" format:"date-time" example:"2021-01-01T00:00:00Z"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at" doc:"Update date of bibliometric" format:"date-time" example:"2021-01-01T00:00:00Z"`
}
type CreateBibliometric struct {
	ID uuid.UUID `db:"id" json:"id" doc:"ID of created Bibliometric" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
}
type CreateBibliometricResponse struct {
	Body CreateBibliometric `json:"body" doc:"Bibliometric creation object"`
}
