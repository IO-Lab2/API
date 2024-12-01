package services

import (
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

func GetBibliometricByAuthor(id uuid.UUID) ([]responses.BibliometricBody, error) {
	bibliometric, err := repositories.BibliometricByAuthor(database.GetDB(), id)
	if err != nil {
		zap.L().Error("Error querying bibliometric by author", zap.Error(err))
		return nil, err
	}

	if len(bibliometric) == 0 {
		zap.L().Warn("No bibliometric found for scientist ID: ", zap.String("scientist_id", id.String()))
		return nil, ErrBibliometricNotFound
	}

	return bibliometric, nil
}
func CreateBibliometric(input *requests.CreateBibliometric) (uuid.UUID, error) {
	id := uuid.New()
	err := repositories.CreateBibliometric(database.GetDB(), id, input)
	if err != nil {
		logging.Logger.Error("Error creating bibliometric", err)
		return uuid.Nil, err
	}
	return id, nil
}
