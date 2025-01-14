package services

import (
	"io-project-api/internal/database"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/repositories"
	"io-project-api/internal/responses"

	"github.com/google/uuid"
)

func GetOrganizationTree(parentID uuid.UUID) (*responses.ListOfOrganizations, error) {

	result, err := repositories.QueryGetOrganizationTree(database.GetDB(), &parentID)
	if len(result.Body) == 0 || err != nil {
		logging.Logger.Error("Error while getting organization tree: ", err)
		return nil, err
	}

	return result, nil
}
