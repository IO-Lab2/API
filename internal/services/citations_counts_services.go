package services

import (
	"errors"
	"io-project-api/internal/database"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/models"
	"io-project-api/internal/repositories"
)

var (
	ErrCitationCountFilterNotFound = errors.New("no citation count found")
)

func GetCitationCount() (*models.CitationsFilter, error) {
	logging.Logger.Info("INFO: Retrieving citation counts")
	db := database.GetDB()

	counts, err := repositories.CitationCountFilter(db)
	if err != nil {
		logging.Logger.Error("ERROR: Error retrieving citation counts:", err)
		return nil, err
	}
	logging.Logger.Info("INFO: Successfully retrieved citation counts")
	return counts, nil
}
