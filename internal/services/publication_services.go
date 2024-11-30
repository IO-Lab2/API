package services

import (
	"database/sql"
	"errors"
	"io-project-api/internal/database"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/repositories"
	"io-project-api/internal/responses"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	ErrPublicationNotFound = errors.New("publication not found for the given ID")
)

func GetPublicationByID(id uuid.UUID) (*responses.PublicationBody, error) {
	publication, err := repositories.PublicationByID(database.GetDB(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			logging.Logger.Warn("No Publication found: ", zap.String("ID", id.String()))
			return nil, ErrPublicationNotFound
		}
		logging.Logger.Error("Error querying Publication by ID: ", zap.Error(err))
		return nil, err
	}

	return publication, nil
}
