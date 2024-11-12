package repositories

import (
	"database/sql"
	"io-project-api/internal/responses"

	"github.com/google/uuid"
)

func PublicationByID(db *sql.DB, id uuid.UUID) ([]responses.PublicationBody, error) {
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
