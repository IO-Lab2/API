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
	ErrScientistPublicationNotFound = errors.New("scientist publication relationship not found for the given ID")
)

func GetScientistPublicationByID(id uuid.UUID) ([]responses.ScientistPublicationBody, error) {
	scientistPublication, err := repositories.ScientistPublicationByID(database.GetDB(), id)
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
