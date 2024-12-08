package services

import (
	"database/sql"
	"errors"
	"io-project-api/internal/database"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/repositories"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	ErrPublicationNotFound = errors.New("publication not found for the given ID")
)

func GetPublicationByID(id uuid.UUID) (*responses.PublicationBody, error) {
	logging.Logger.Info("INFO: Retrieving publication by ID")
	publication, err := repositories.PublicationByID(database.GetDB(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			logging.Logger.Warn("WARN: No Publication found: ", zap.String("ID", id.String()))
			return nil, ErrPublicationNotFound
		}
		logging.Logger.Error("ERROR: Error querying Publication by ID: ", zap.Error(err))
		return nil, err
	}
	logging.Logger.Info("INFO: Successfully retrieving publication by ID")

	return publication, nil
}
func CreatePublication(input *requests.CreatePublicationRequest) (uuid.UUID, error) {
	id := uuid.New()
	logging.Logger.Info("INFO: Creating Publication")
	err := repositories.CreatePublication(database.GetDB(), id, input)
	if err != nil {
		logging.Logger.Error("ERROR: Error creating publication", zap.Error(err))
		return uuid.Nil, err
	}
	logging.Logger.Info("INFO: Successfully creating Publication")
	return id, nil
}
