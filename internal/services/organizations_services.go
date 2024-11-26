package services

import (
	"errors"
	"io-project-api/internal/database"
	"io-project-api/internal/repositories"
	"io-project-api/internal/responses"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	ErrOrganizationNotFound = errors.New("organization not found for the given ID")
)

func GetOrganizationByID(id uuid.UUID) (*responses.OrganizationBody, error) {
	organization, err := repositories.OrganizationByID(database.GetDB(), id)
	if err != nil {
		zap.L().Error("Error querying organization by ID", zap.Error(err))
		return nil, err
	}

	return organization, nil
}

func GetOrganizationsByScientistID(id uuid.UUID) ([]responses.OrganizationBody, error) {
	organizations, err := repositories.OrganizationsByScientistID(database.GetDB(), id)
	if err != nil {
		zap.L().Error("Error querying organizations by scientist ID", zap.Error(err))
		return nil, err
	}

	return organizations, nil
}
