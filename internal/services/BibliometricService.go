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
	ErrBibliometricNotFound = errors.New("bibliometric not found for the given ID")
)

func GetBibliometricByID(id uuid.UUID) (*responses.BibliometricBody, error) {
	bibliometric, err := repositories.BibliometricByID(database.GetDB(), id)
	if err != nil {
		zap.L().Error("Error querying bibliometric by ID", zap.Error(err))
		return nil, err
	}

	return bibliometric, nil
}

func GetBibliometricByAuthor(author string) ([]responses.BibliometricBody, error) {
	bibliometric, err := repositories.BibliometricByAuthor(database.GetDB(), author)
	if err != nil {
		zap.L().Error("Error querying bibliometric by author", zap.Error(err))
		return nil, err
	}

	if len(bibliometric) == 0 {
		zap.L().Warn("No bibliometric found", zap.String("Author", author))
		return nil, ErrBibliometricNotFound
	}

	return bibliometric, nil
}
