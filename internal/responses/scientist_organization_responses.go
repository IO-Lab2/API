package responses

import (
	"time"

	"github.com/google/uuid"
)

type ScientistOrganizationResponse struct {
	Body []ScientistOrganizationBody `json:"body" doc:"Scientist Organization response body"`
}

type ScientistOrganizationBody struct {
	ID             uuid.UUID `json:"id" doc:"Scientist Organization ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
	ScientistID    uuid.UUID `json:"scientist_id" doc:"Scientist ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
	OrganizationID uuid.UUID `json:"organization_id" doc:"Organization ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
	CreatedAt      time.Time `json:"created_at" doc:"Creation date of the scientist organization" format:"date-time" example:"2021-01-01T00:00:00Z"`
	UpdatedAt      time.Time `json:"updated_at" doc:"Last update date of the scientist organization" format:"date-time" example:"2021-01-01T00:00:00Z"`
}
