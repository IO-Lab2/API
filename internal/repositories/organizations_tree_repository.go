package repositories

import (
	"io-project-api/internal/responses"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func QueryGetOrganizationTree(db *sqlx.DB, parentID *uuid.UUID) (*responses.ListOfOrganizations, error) {

	if parentID == nil || *parentID == uuid.Nil {
		// Query to get root organizations
		query := `
			SELECT o.id, o.name, o.type
			FROM organizations o
			INNER JOIN organizations_relationships orl
			ON o.id = orl.child_id
			WHERE orl.parent_id IS NULL`

		// Execute query
		result := responses.ListOfOrganizations{}
		if err := db.Select(&result.Body, query); err != nil {
			return nil, err
		}
		return &result, nil
	}

	query := `
		SELECT o.id, o.name, o.type
		FROM organizations o
		INNER JOIN organizations_relationships orl
		ON o.id = orl.child_id
		WHERE orl.parent_id = $1`

	// Execute query
	result := responses.ListOfOrganizations{}
	if err := db.Select(&result.Body, query, parentID); err != nil {
		return nil, err
	}

	return &result, nil
}
