package services

import (
	"errors"
	"io-project-api/internal/database"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/models"
	"io-project-api/internal/repositories"
)

var (
	ErrMinisterialScoreFilterNotFound = errors.New("no ministerial scores found")
)

func GetMinisterialScores() (*models.MinisterialScore, error) {
	logging.Logger.Info("INFO: Retrieving ministerial counts")
	db := database.GetDB()

	scores, err := repositories.MinisterialScoreFilter(db)
	if err != nil {
		logging.Logger.Error("ERROR: Error retrieving ministerial scores ", err)
		return nil, err
	}
	logging.Logger.Info("INFO: Successfully retrieving ministerial counts")

	return scores, nil
}
