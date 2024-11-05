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
	ErrPublicationNotFound = errors.New("publication not found for the given ID")
)

func GetPublicationByID(db *sql.DB, id uuid.UUID) ([]models.Publication, error) {
	publication, err := repositories.PublicationByID(db, id)
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
