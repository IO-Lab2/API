package services

import (
	"errors"
	"io-project-api/internal/database"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/models"
	"io-project-api/internal/repositories"

	"go.uber.org/zap"
)

var (
	ErrPublicationCountFilterNotFound = errors.New("no publication count found")
)

func GetPublicationCount() (*models.PublicationCount, error) {
	logging.Logger.Info("INFO: Retrieving publication count")
	db := database.GetDB()

	counts, err := repositories.PublicationCountFilter(db)
	if err != nil {
		logging.Logger.Error("ERROR: Error retrieving publication counts", zap.Error(err))
		return nil, err
	}
	logging.Logger.Info("INFO: Successfully retrieving publication count")
	return counts, nil
}
