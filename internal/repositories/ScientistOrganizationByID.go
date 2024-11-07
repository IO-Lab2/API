package repositories

import (
	"database/sql"
	"io-project-api/internal/models"

	"github.com/google/uuid"
)

func ScientistOragnizationByID(db *sql.DB, id uuid.UUID) ([]models.ScientistOrganization, error) {
	query := "SELECT id, scientist_id, organization_id FROM scientist_organization WHERE id = $1"
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}

	var scientistsOrganizations []models.ScientistOrganization

	for rows.Next() {
		var scientistOrganization models.ScientistOrganization
		err := rows.Scan(
			&scientistOrganization.ID,
			&scientistOrganization.ScientistID,
			&scientistOrganization.OrganizationID,
			&scientistOrganization.CreatedAt,
			&scientistOrganization.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		scientistsOrganizations = append(scientistsOrganizations, scientistOrganization)
	}
	return scientistsOrganizations, nil
}
