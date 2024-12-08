package requests

import (
	"time"

	"github.com/google/uuid"
)

type PublicationID struct {
	ID uuid.UUID `path:"id" doc:"Publication ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
}
type CreatePublicationRequest struct {
	Title               string    `query:"title" json:"title" doc:"Publication title" format:"string" example:"Publication title"`
	Journal             string    `query:"journal" json:"journal" doc:"Journal name" format:"string" example:"Journal name"`
	PublicationDate     time.Time `query:"publication_date" json:"publication_date" doc:"Publication date" format:"date-time" example:"2021-01-01T15:04:05Z"`
	JournalImpactFactor float64   `query:"journal_impact_factor" json:"journal_impact_factor" doc:"Journal impact factor" format:"float" example:"1.23"`
	Publisher           string    `query:"publisher" json:"publisher" doc:"Publisher name" format:"string" example:"Publisher name"`
	JournalType         string    `query:"journal_type" json:"journal_type" doc:"Journal type" format:"string" example:"Journal type"`
}
type DeletePublication struct {
	ID uuid.UUID `path:"id" doc:"Publication ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
}
