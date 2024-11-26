package repositories

import (
	"io-project-api/internal/models"
	"io-project-api/internal/responses"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func PublicationByID(db *sqlx.DB, id uuid.UUID) ([]responses.PublicationBody, error) {
	query := "SELECT id, title, journal, publication_date, citations_count, impact_factor, scientist_id, created_at,  updated_at FROM publications WHERE id =$1"
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}

	var Publications []responses.PublicationBody

	for rows.Next() {
		var Publication responses.PublicationBody
		err := rows.Scan(
			&Publication.ID,
			&Publication.Title,
			&Publication.Journal,
			&Publication.PublicationDate,
			&Publication.CitationsCount,
			&Publication.ImpactFactor,
			&Publication.CreatedAt,
			&Publication.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		Publications = append(Publications, Publication)
	}
	return Publications, nil
}

func PublicationsByScientistID(db *sqlx.DB, id uuid.UUID) ([]responses.PublicationBody, error) {
	query := `
		SELECT p.*
		FROM publications p
		JOIN scientist_publication sp ON p.id = sp.publication_id
		WHERE sp.scientist_id = $1`
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}

	var Publications []responses.PublicationBody

	for rows.Next() {
		var Publication responses.PublicationBody
		err := rows.Scan(
			&Publication.ID,
			&Publication.Title,
			&Publication.Journal,
			&Publication.PublicationDate,
			&Publication.CitationsCount,
			&Publication.ImpactFactor,
			&Publication.CreatedAt,
			&Publication.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		Publications = append(Publications, Publication)
	}
	return Publications, nil
}

func PublicationCountFilter(db *sqlx.DB) (*models.PublicationCount, error) {
	query := "SELECT MAX(citations_count) as largest, MIN(citations_count) as smallest FROM publications"
	var publicationCount models.PublicationCount
	if err := db.Get(&publicationCount, query); err != nil {
		return nil, err
	}

	return &publicationCount, nil
}
