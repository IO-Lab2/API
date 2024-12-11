package repositories

import (
	logging "io-project-api/internal/logger"
	"io-project-api/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Query which returns all scientist's publication by scientist id
func ScientistPublicationByID(db *sqlx.DB, id uuid.UUID) ([]models.Publication, error) {
	query := `
		SELECT p.id, p.title, p.journal, p.publication_date, p.journal_impact_factor, p.journal_type, p.publisher, p.created_at, p.updated_at
		FROM publications p
		JOIN scientists_publications sp ON p.id = sp.publication_id
		WHERE sp.scientist_id = $1`

	var publications []models.Publication
	if err := db.Select(&publications, query, id); err != nil {
		logging.Logger.Error("ERROR: Error querying scientist-publication relationships by ID", zap.Error(err))
		return nil, err
	}

	logging.Logger.Info("INFO: Successfully retrieved scientist-publication relationships by ID")
	return publications, nil
}
