package repositories

import (
	"io-project-api/internal/responses"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func ScientistOragnizationByID(db *sqlx.DB, id uuid.UUID) ([]responses.ScientistOrganizationBody, error) {
	query := "SELECT id, scientist_id, organization_id FROM scientist_organization WHERE id = $1"
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}

	var scientistsOrganizations []responses.ScientistOrganizationBody

	for rows.Next() {
		var scientistOrganization responses.ScientistOrganizationBody
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
