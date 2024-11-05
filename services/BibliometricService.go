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
	ErrBibliometricNotFound = errors.New("bibliometric not found for the given ID")
)

func GetBibliometricByID(db *sql.DB, id uuid.UUID) ([]models.Bibliometrics, error) {
	bibliometric, err := repositories.BibliometricByID(db, id)
	if err != nil {
		zap.L().Error("Error querying bibliometric by ID", zap.Error(err))
		return nil, err
	}

	if len(bibliometric) == 0 {
		zap.L().Warn("No bibliometric found", zap.String("ID", id.String()))
		return nil, ErrBibliometricNotFound
	}

	return bibliometric, nil
}
