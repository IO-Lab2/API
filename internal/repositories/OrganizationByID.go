package repositories

import (
	"io-project-api/internal/responses"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func OrganizationByID(db *sqlx.DB, id uuid.UUID) ([]responses.OrganizationBody, error) {
	query := "SELECT id, name, type organization FROM organizations WHERE id = $1"
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}

	var Organizations []responses.OrganizationBody

	for rows.Next() {
		var Organization responses.OrganizationBody
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
