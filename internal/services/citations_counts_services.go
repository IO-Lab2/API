package services

import (
	"errors"
	"io-project-api/internal/database"
	"io-project-api/internal/models"
	"io-project-api/internal/repositories"

	"go.uber.org/zap"
)

var (
	ErrCitationCountFilterNotFound = errors.New("no citation count found")
)

func GetCitationCount() (*models.CitationsFilter, error) {

	db := database.GetDB()

	counts, err := repositories.CitationCountFilter(db)
	if err != nil {
		zap.L().Error("Error retrieving citation counts", zap.Error(err))
		return nil, err
	}

	return counts, nil
}
