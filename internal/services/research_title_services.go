package services

import (
	"errors"
	"io-project-api/internal/database"
	"io-project-api/internal/models"
	"io-project-api/internal/repositories"

	"go.uber.org/zap"
)

var (
	ErrResearchTitleFilterNotFound = errors.New("no research title found")
)

func GetResearchAreas() ([]models.ResearchArea, error) {

	db := database.GetDB()

	titles, err := repositories.ResearchAreaFilter(db)
	if err != nil {
		zap.L().Error("Error retrieving research titles", zap.Error(err))
		return nil, err
	}

	if len(titles) == 0 {
		zap.L().Warn("No research titles found in database")
		return nil, ErrResearchTitleFilterNotFound
	}

	return titles, nil
}
