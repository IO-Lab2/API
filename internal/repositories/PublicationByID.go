package repositories

import (
	"database/sql"
	"io-project-api/internal/models"

	"github.com/google/uuid"
)

func PublicationByID(db *sql.DB, id uuid.UUID) ([]models.Publication, error) {
	query := "SELECT id, title, journal, publication_date, citations_count, impact_factor, scientist_id, created_at,  updated_at FROM publications WHERE id =$1"
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}

	var Publications []models.Publication

	for rows.Next() {
		var Publication models.Publication
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
