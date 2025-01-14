package handlers

import (
	"context"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetOrganizationTreeHandler(ctx context.Context, input *requests.OrganizationTreeFilterRequest) (*responses.ListOfOrganizations, error) {

	result, err := services.GetOrganizationTree(input.ParentID)
	if result == nil || err != nil {
		logging.Logger.Error("Error while getting organization tree: ", err)
		return nil, huma.Error400BadRequest("Error: No organization found")
	}
	return result, nil
}
