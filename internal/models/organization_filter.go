package models

type OrganizationFilter struct {
	OrganizationName string `db:"organization" json:"organization"`
}
