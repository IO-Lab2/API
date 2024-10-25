package models

import (
	"database/sql/driver"
	"errors"
)

type OrganizationEnum string

const (
	University OrganizationEnum = "university"
	Institute  OrganizationEnum = "institute"
	Cathedra   OrganizationEnum = "cathedra"
)

func (org OrganizationEnum) Value() (driver.Value, error) {
	return string(org), nil
}
func (org *OrganizationEnum) Scan(value interface{}) error {
	if v, ok := value.(string); ok {
		*org = OrganizationEnum(v)
		return nil
	}
	return errors.New("OrganizationEnum scan failed")
}
func (org OrganizationEnum) String() string {
	return string(org)
}
