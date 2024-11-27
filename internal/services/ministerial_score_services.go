package services

import (
	"errors"
	"io-project-api/internal/database"
	"io-project-api/internal/models"
	"io-project-api/internal/repositories"

	"go.uber.org/zap"
)

var (
	ErrMinisterialScoreFilterNotFound = errors.New("no ministerial scores found")
)

func GetMinisterialScores() (*models.MinisterialScore, error) {

	db := database.GetDB()

	scores, err := repositories.MinisterialScoreFilter(db)
	if err != nil {
		zap.L().Error("Error retrieving ministerial scores: ", zap.Error(err))
		return nil, err
	}

	return scores, nil
}
