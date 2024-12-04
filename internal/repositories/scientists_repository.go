package repositories

import (
	logging "io-project-api/internal/logger"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func ScientistByID(db *sqlx.DB, id uuid.UUID) ([]responses.ScientistBody, error) {
	query := "SELECT id, first_name, last_name, academic_title, research_area, email, profile_url, created_at, updated_at FROM scientists WHERE id = $1"
	logging.Logger.Info("INFO: Executing query:", query)

	rows, err := db.Query(query, id)
	if err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
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
			logging.Logger.Error("ERROR: Error scanning row:", err)
			return nil, err
		}
		scientists = append(scientists, scientist)
	}

	if err := rows.Err(); err != nil {
		logging.Logger.Error("ERROR: Error iterating over rows:", err)
		return nil, err
	}

	logging.Logger.Info("INFO: Successfully retrieved scientists by ID")
	return scientists, nil
}

func ScientistByName(db *sqlx.DB, name *requests.ScientistName) (*responses.ScientistBody, error) {
	query := "SELECT id, first_name, last_name, academic_title, research_area, email, profile_url, created_at, updated_at FROM scientists WHERE first_name = $1 AND last_name = $2"
	logging.Logger.Info("INFO: Executing query:", query)

	var scientist responses.ScientistBody
	if err := db.Get(&scientist, query, name.FirstName, name.LastName); err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}

	logging.Logger.Info("INFO: Successfully retrieved scientist by name")
	return &scientist, nil
}
