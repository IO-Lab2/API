package requests

import "github.com/google/uuid"

type OrganizationTreeFilterRequest struct {
	ParentID uuid.UUID `query:"id" doc:"The ID of the parent organization. If not provided, the root organizations will be returned." format:"uuid" example:"00000000-0000-0000-0000-000000000000"`
}
