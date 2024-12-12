package services

import (
	"errors"
	"io-project-api/internal/database"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/repositories"
	"io-project-api/internal/responses"

	"go.uber.org/zap"
)

var (
	ErrResearchTitleFilterNotFound = errors.New("no research title found")
)

func GetResearchAreas() ([]responses.ResearchAreaExtended, error) {
	logging.Logger.Info("INFO: Retrieving research areas")

	db := database.GetDB()

	titles, err := repositories.ResearchAreaFilter(db)
	if err != nil {
		logging.Logger.Error("ERROR: Error retrieving research titles", zap.Error(err))
		return nil, err
	}

	if len(titles) == 0 {
		logging.Logger.Warn("No research titles found in database")
		return nil, ErrResearchTitleFilterNotFound
	}
	logging.Logger.Info("INFO: Successfully retrieving research areas")
	return titles, nil
}
