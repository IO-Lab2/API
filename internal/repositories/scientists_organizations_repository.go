package repositories

import (
	logging "io-project-api/internal/logger"
	"io-project-api/internal/responses"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func ScientistOrganizationByID(db *sqlx.DB, id uuid.UUID) ([]responses.ScientistOrganizationBody, error) {
	query := "SELECT id, scientist_id, organization_id, created_at, updated_at FROM scientist_organization WHERE id = $1"
	logging.Logger.Info("INFO: Executing query:", query)

	rows, err := db.Query(query, id)
	if err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

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
			logging.Logger.Error("ERROR: Error scanning row:", err)
			return nil, err
		}
		scientistsOrganizations = append(scientistsOrganizations, scientistOrganization)
	}

	if err := rows.Err(); err != nil {
		logging.Logger.Error("ERROR: Error iterating over rows:", err)
		return nil, err
	}

	logging.Logger.Info("INFO: Successfully retrieved scientist-organization relationships by ID")
	return scientistsOrganizations, nil
}

func ScientistsOrganizationByScientistID(db *sqlx.DB, id uuid.UUID) ([]responses.ScientistOrganizationBody, error) {
	query := `
		SELECT id, scientist_id, organization_id, created_at, updated_at
		FROM scientist_organization
		WHERE scientist_id = $1`
	logging.Logger.Info("INFO: Executing query:", query)

	rows, err := db.Query(query, id)
	if err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

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
			logging.Logger.Error("ERROR: Error scanning row:", err)
			return nil, err
		}
		scientistsOrganizations = append(scientistsOrganizations, scientistOrganization)
	}

	if err := rows.Err(); err != nil {
		logging.Logger.Error("ERROR: Error iterating over rows:", err)
		return nil, err
	}

	logging.Logger.Info("INFO: Successfully retrieved scientist-organization relationships by scientist ID")
	return scientistsOrganizations, nil
}
