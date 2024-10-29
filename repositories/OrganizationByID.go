package models

import (
	"database/sql"
	"io-project-api/models"

	"github.com/google/uuid"
)

func OrganizationByID(db *sql.DB, id uuid.UUID) ([]models.Organization, error) {
	query := "SELECT id, name, type organization FROM organizations WHERE id = $1"
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}

	var Organizations []models.Organization

	for rows.Next() {
		var Organization models.Organization
		err := rows.Scan(
			&Organization.ID,
			&Organization.Name,
			&Organization.Type,
			&Organization.CreatedAt,
			&Organization.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		Organizations = append(Organizations, Organization)
	}
	return Organizations, nil
}
