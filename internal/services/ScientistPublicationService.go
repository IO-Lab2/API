package services

import (
	"database/sql"
	"errors"
	"io-project-api/internal/models"
	"io-project-api/internal/repositories"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	ErrScientistPublicationNotFound = errors.New("scientist publication relationship not found for the given ID")
)

func GetScientistPublicationByID(db *sql.DB, id uuid.UUID) ([]models.ScientistPublication, error) {
	scientistPublication, err := repositories.ScientistPublicationByID(db, id)
	if err != nil {
		zap.L().Error("Error querying scientist publication relationship by id", zap.Error(err))
		return nil, err
	}
	if len(scientistPublication) == 0 {
		zap.L().Error("No scientist publication relationship found", zap.String("ID", id.String()))
		return nil, ErrScientistOrganizationNotFound
	}
	return scientistPublication, nil
}
