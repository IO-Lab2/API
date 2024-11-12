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
	ErrPublicationNotFound = errors.New("publication not found for the given ID")
)

func GetPublicationByID(id uuid.UUID) ([]responses.PublicationBody, error) {
	publication, err := repositories.PublicationByID(database.GetDB().DB, id)
	if err != nil {
		zap.L().Error("Error querying Publication by ID", zap.Error(err))
		return nil, err
	}

	if len(publication) == 0 {
		zap.L().Warn("No Publication found", zap.String("ID", id.String()))
		return nil, ErrPublicationNotFound
	}

	return publication, nil
}
