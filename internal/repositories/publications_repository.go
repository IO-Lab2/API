package repositories

import (
	"io-project-api/internal/models"
	"io-project-api/internal/responses"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// PublicationByID retrieves a publication entry by its ID.
func PublicationByID(db *sqlx.DB, id uuid.UUID) (*responses.PublicationBody, error) {
	query := "SELECT id, title, journal, publication_date, journal_impact_factor, created_at, updated_at FROM publications WHERE id = $1"
	var publication responses.PublicationBody
	if err := db.Get(&publication, query, id); err != nil {
		return nil, err
	}
	return &publication, nil
}

// PublicationsByScientistID retrieves publication entries by the scientist's ID.
func PublicationsByScientistID(db *sqlx.DB, id uuid.UUID) ([]responses.PublicationBody, error) {
	query := `
		SELECT p.id, p.title, p.journal, p.publication_date, p.journal_impact_factor, p.created_at, p.updated_at
		FROM publications p
		JOIN scientist_publication sp ON p.id = sp.publication_id
		WHERE sp.scientist_id = $1`
	var publications []responses.PublicationBody
	if err := db.Select(&publications, query, id); err != nil {
		return nil, err
	}
	return publications, nil
}

func PublicationCountFilter(db *sqlx.DB) (*models.PublicationCount, error) {
	query := "SELECT MAX(citations_count) as largest, MIN(citations_count) as smallest FROM publications"
	var publicationCount models.PublicationCount
	if err := db.Get(&publicationCount, query); err != nil {
		return nil, err
	}

	return &publicationCount, nil
}
