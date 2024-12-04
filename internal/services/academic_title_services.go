package services

import (
	"errors"
	"io-project-api/internal/database"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/models"
	"io-project-api/internal/repositories"
)

var (
	ErrAcademicTitleFilterNotFound = errors.New("no academic title found")
)

func GetAcademicTitles() ([]models.AcademicTitle, error) {
	logging.Logger.Info("INFO: Retrieving Academic titles")
	db := database.GetDB()

	titles, err := repositories.AcademicTitleFilter(db)
	if err != nil {
		logging.Logger.Error("ERROR: Error retrieving academic titles:", err)
		return nil, err
	}

	if len(titles) == 0 {
		logging.Logger.Warn("Warn: No academic titles found in database")
		return nil, ErrAcademicTitleFilterNotFound
	}
	logging.Logger.Info("INFO: Successfully retrieved academic titles")
	return titles, nil
}
