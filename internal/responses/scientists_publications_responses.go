package responses

import (
	"time"

	"github.com/google/uuid"
)

type ScientistPublicationResponse struct {
	Body []ScientistPublicationBody `json:"body" doc:"Scientist publication response body"`
}

type ScientistPublicationBody struct {
	ID            uuid.UUID `json:"id" doc:"Scientist publication ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
	ScientistID   uuid.UUID `json:"scientist_id" doc:"Scientist ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
	PublicationID uuid.UUID `json:"publication_id" doc:"Publication ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
	CreatedAt     time.Time `json:"created_at" doc:"Creation date of ScientistPublication" format:"date-time" example:"2021-01-01T00:00:00Z"`
	UpdatedAt     time.Time `json:"updated_at" doc:"Date of last update of ScientistPublication" format:"date-time" example:"2021-01-01T00:00:00Z"`
}
