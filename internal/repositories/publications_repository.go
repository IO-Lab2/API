package repositories

import (
	logging "io-project-api/internal/logger"
	"io-project-api/internal/models"
	"io-project-api/internal/responses"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func PublicationByID(db *sqlx.DB, id uuid.UUID) (*responses.PublicationBody, error) {
	query := "SELECT id, title, journal, publication_date, journal_impact_factor, created_at, updated_at FROM publications WHERE id = $1"
	logging.Logger.Info("INFO: Executing query:", query)

	var publication responses.PublicationBody
	if err := db.Get(&publication, query, id); err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}
	logging.Logger.Info("INFO: Successfully retrieved publication by ID")
	return &publication, nil
}

func PublicationsByScientistID(db *sqlx.DB, id uuid.UUID) ([]responses.PublicationBody, error) {
	query := `
		SELECT p.id, p.title, p.journal, p.publication_date, p.journal_impact_factor, p.created_at, p.updated_at
		FROM publications p
		JOIN scientist_publication sp ON p.id = sp.publication_id
		WHERE sp.scientist_id = $1`
	logging.Logger.Info("INFO: Executing query:", query)

	var publications []responses.PublicationBody
	if err := db.Select(&publications, query, id); err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}
	logging.Logger.Info("INFO: Successfully retrieved publications by scientist ID")
	return publications, nil
}

func PublicationCountFilter(db *sqlx.DB) (*models.PublicationCount, error) {
	query := "SELECT MAX(citations_count) as largest, MIN(citations_count) as smallest FROM publications"
	logging.Logger.Info("INFO: Executing query:", query)

	var publicationCount models.PublicationCount
	if err := db.Get(&publicationCount, query); err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}
	logging.Logger.Info("INFO: Successfully retrieved publication count filter")
	return &publicationCount, nil
}
