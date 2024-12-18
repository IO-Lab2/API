package repositories

import (
	"io-project-api/internal/models"

	"github.com/jmoiron/sqlx"
)

func QueryGetJournalTypes(db *sqlx.DB) ([]models.JournalType, error) {
	query := `
		SELECT DISTINCT journal_type 
		FROM publications
		WHERE journal_type IS NOT NULL`
	var journalTypes []models.JournalType
	err := db.Select(&journalTypes, query)
	if err != nil {
		return nil, err
	}
	return journalTypes, nil
}
