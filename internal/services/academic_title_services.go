package services

import (
	"errors"
	"io-project-api/internal/database"
	"io-project-api/internal/models"
	"io-project-api/internal/repositories"

	"go.uber.org/zap"
)

var (
	ErrAcademicTitleFilterNotFound = errors.New("no academic title found")
)

func GetAcademicTitles() ([]models.AcademicTitle, error) {

	db := database.GetDB()

	titles, err := repositories.AcademicTitleFilter(db)
	if err != nil {
		zap.L().Error("Error retrieving academic titles", zap.Error(err))
		return nil, err
	}

	if len(titles) == 0 {
		zap.L().Warn("No academic titles found in database")
		return nil, ErrAcademicTitleFilterNotFound
	}

	return titles, nil
}
