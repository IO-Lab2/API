package services

import (
	"errors"
	"io-project-api/internal/repositories"
	"io-project-api/internal/database"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	ErrMinisterialScoreFilterNotFound = errors.New("no ministerial scores found")
)

func GetMinisterialScores() ([]models.MinisterialScore, error) {
	
	db := database.GetDB()

	
	scores, err := repositories.MinisterialScoreFilter(db)
	if err != nil {
		zap.L().Error("Error retrieving ministerial scores", zap.Error(err))
		return nil, err
	}

	
	if len(scores) == 0 {
		zap.L().Warn("No ministerial scores found in database")
		return nil, ErrNoMinisterialScores
	}

	
	return scores, nil
}

