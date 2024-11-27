package repositories

import (
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func ScientistByID(db *sqlx.DB, id uuid.UUID) ([]responses.ScientistBody, error) {
	query := "SELECT id, first_name, last_name, academic_title, research_area, email, profile_url, created_at, updated_at FROM scientists WHERE id = $1"
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scientists []responses.ScientistBody

	for rows.Next() {
		var scientist responses.ScientistBody
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

func ScientistByName(db *sqlx.DB, name *requests.ScientistName) (*responses.ScientistBody, error) {
	query := "SELECT id, first_name, last_name, academic_title, research_area, email, profile_url, created_at, updated_at FROM scientists WHERE first_name = $1 AND last_name = $2"
	var scientist *responses.ScientistBody
	if err := db.Get(&scientist, query, name.FirstName, name.LastName); err != nil {
		return nil, err
	}

	return scientist, nil
}
