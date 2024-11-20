package services

import (
	"errors"
	"io-project-api/internal/repositories"
	"io-project-api/internal/database"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	ErrCitationCountFilterNotFound = errors.New("no citation count found")
)

func GetCitationCount() ([]models.CitationCount, error) {
	
	db := database.GetDB()

	
	counts, err := repositories.CitationCountFilter(db)
	if err != nil {
		zap.L().Error("Error retrieving citation counts", zap.Error(err))
		return nil, err
	}

	
	if len(counts) == 0 {
		zap.L().Warn("No citation counts found in database")
		return nil, ErrCitationCountFilterNotFound
	}

	
	return counts, nil
}

