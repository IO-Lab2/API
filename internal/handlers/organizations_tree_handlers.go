package handlers

import (
	"context"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"
)

func GetOrganizationTreeHandler(ctx context.Context, input *requests.OrganizationTreeFilterRequest) (*responses.ListOfOrganizations, error) {

	result, err := services.GetOrganizationTree(input.ParentID)
	if err != nil {
		logging.Logger.Error("Error while getting organization tree: ", err)
		return nil, err
	}
	return result, nil
}
