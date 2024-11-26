package services

import (
	"errors"
	"io-project-api/internal/database"
	"io-project-api/internal/models"
	"io-project-api/internal/repositories"

	"go.uber.org/zap"
)

var (
	ErrPublicationCountFilterNotFound = errors.New("no publication count found")
)

func GetPublicationCount() ([]models.PublicationCount, error) {

	db := database.GetDB()

	counts, err := repositories.PublicationCountFilter(db)
	if err != nil {
		zap.L().Error("Error retrieving publication counts", zap.Error(err))
		return nil, err
	}

	if len(counts) == 0 {
		zap.L().Warn("No publication counts found in database")
		return nil, ErrPublicationCountFilterNotFound
	}

	return counts, nil
}
