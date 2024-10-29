package models

import (
	"database/sql"
	"io-project-api/models"

	"github.com/google/uuid"
)

func ScientistByID(db *sql.DB, id uuid.UUID) ([]models.Scientist, error) {
	query := "SELECT id first_name, last_name,  academic_title,  research_area, email,  profile_url FROM scientists WHERE id = $1"
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}

	var scientists []models.Scientist

	for rows.Next() {
		var scientist models.Scientist
		err := rows.Scan(
			&scientist.ID,
			&scientist.FirstName,
			&scientist.LastName,
			&scientist.AcademicTitle,
			&scientist.ResearchArea,
			&scientist.Email,
			&scientist.ProfileUrl,
			&scientist.CreatedAt,
			&scientist.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		scientists = append(scientists, scientist)
	}
	return scientists, nil
}
