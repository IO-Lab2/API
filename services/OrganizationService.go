package services

import (
	"database/sql"
	"errors"
	"io-project-api/models"
	"io-project-api/repositories"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	ErrOrganizationNotFound = errors.New("organization not found for the given ID")
)

func GetOrganizationByID(db *sql.DB, id uuid.UUID) ([]models.Organization, error) {
	organization, err := repositories.OrganizationByID(db, id)
	if err != nil {
		zap.L().Error("Error querying organization by ID", zap.Error(err))
		return nil, err
	}

	if len(organization) == 0 {
		zap.L().Warn("No organization found", zap.String("ID", id.String()))
		return nil, ErrOrganizationNotFound
	}

	return organization, nil
}
