package repositories

import (
	logging "io-project-api/internal/logger"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func OrganizationByID(db *sqlx.DB, id uuid.UUID) (*responses.OrganizationBodyExtended, error) {
	query := "SELECT id, name, type, created_at, updated_at FROM organizations WHERE id = $1"
	logging.Logger.Info("INFO: Executing query:", query)

	var organization responses.OrganizationBodyExtended
	if err := db.Get(&organization, query, id); err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}
	logging.Logger.Info("INFO: Successfully retrieved organization by ID")
	return &organization, nil
}

func OrganizationsByScientistID(db *sqlx.DB, id uuid.UUID) ([]responses.OrganizationBodyExtended, error) {
	query := `
		SELECT o.id, o.name, o.type, o.created_at, o.updated_at
		FROM organizations o
		JOIN scientist_organization so ON o.id = so.organization_id
		WHERE so.scientist_id = $1`
	logging.Logger.Info("INFO: Executing query:", query)

	rows, err := db.Query(query, id)
	if err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var organizations []responses.OrganizationBodyExtended
	for rows.Next() {
		var organization responses.OrganizationBodyExtended
		err := rows.Scan(
			&organization.ID,
			&organization.Name,
			&organization.Type,
			&organization.CreatedAt,
			&organization.UpdatedAt,
		)
		if err != nil {
			logging.Logger.Error("ERROR: Error scanning row:", err)
			return nil, err
		}
		organizations = append(organizations, organization)
	}

	if err = rows.Err(); err != nil {
		logging.Logger.Error("ERROR: Error iterating over rows:", err)
		return nil, err
	}

	logging.Logger.Info("INFO: Successfully retrieved organizations by scientist ID")
	return organizations, nil
}

func Organizations(db *sqlx.DB) (*responses.ListOfOrganizations, error) {
	query := "SELECT id, name, type, created_at, updated_at FROM organizations"
	logging.Logger.Info("INFO: Executing query:", query)

	rows, err := db.Query(query)
	if err != nil {
		logging.Logger.Error("ERROR: Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var organizations []responses.OrganizationBody
	for rows.Next() {
		var organization responses.OrganizationBody
		err := rows.Scan(
			&organization.ID,
			&organization.Name,
			&organization.Type,
		)
		if err != nil {
			logging.Logger.Error("ERROR: Error scanning row:", err)
			return nil, err
		}
		organizations = append(organizations, organization)
	}

	if err = rows.Err(); err != nil {
		logging.Logger.Error("ERROR: Error iterating over rows:", err)
		return nil, err
	}

	logging.Logger.Info("INFO: Successfully retrieved all organizations")
	organizationResponse := &responses.ListOfOrganizations{Body: organizations}
	return organizationResponse, nil
}

func CreateOrganization(db *sqlx.DB, id uuid.UUID, input *requests.CreateOrganization) error {
	query := `
        INSERT INTO organizations (id,name, type, created_at, updated_at)
        VALUES ($1, $2, $3, NOW(), NOW())
		`
	_, err := db.Exec(query, id, input.Name, input.OrganizationType)
	return err

}
