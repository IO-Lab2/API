package services

import (
	"errors"
	"io-project-api/internal/database"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/models"
	"io-project-api/internal/repositories"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	ErrScientistPublicationNotFound = errors.New("scientist publication relationship not found for the given ID")
)

func GetScientistPublicationByID(id uuid.UUID) ([]models.Publication, error) {
	logging.Logger.Info("INFO: Retrieving scientist publication by ID")
	scientistPublication, err := repositories.ScientistPublicationByID(database.GetDB(), id)
	if err != nil {
		logging.Logger.Error("ERROR: Error querying scientist publication relationship by id", zap.Error(err))
		return nil, err
	}
	if len(scientistPublication) == 0 {
		logging.Logger.Error("ERROR: No scientist publication relationship found", zap.String("ID", id.String()))
		return nil, ErrScientistOrganizationNotFound
	}
	logging.Logger.Info("INFO: Successfully retrieving scientist publication by ID")
	return scientistPublication, nil
}
