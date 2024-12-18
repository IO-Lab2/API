package services

import (
	"io-project-api/internal/database"
	"io-project-api/internal/models"
	"io-project-api/internal/repositories"
)

func GetJournalTypes() ([]models.JournalType, error) {
	journalTypes, err := repositories.QueryGetJournalTypes(database.GetDB())
	if err != nil {
		return nil, err
	}
	return journalTypes, nil
}
