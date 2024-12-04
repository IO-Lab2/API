package requests

import (
	"io-project-api/internal/models"

	"github.com/google/uuid"
)

type OrganizationID struct {
	ID uuid.UUID `path:"id" doc:"Organization ID" format:"uuid" example:"8c4bfb01-3c0a-416c-a07c-a24ee52a8b2a"`
}
type CreateOrganization struct {
	Name             string                  `query:"name" json:"name" doc:"Name of the organization" format:"string" example:"Department of Hydrotechnics, Technology and Management"`
	OrganizationType models.OrganizationEnum `query:"type" json:"type" doc:"Type of the organization" format:"string" example:"cathedra"`
}

type OrganizationFilterRequest struct {
}
