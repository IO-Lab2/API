package models

import (
	"database/sql/driver"
	"errors"
)

type Organization string

const (
	University Organization = "university"
	Institute  Organization = "institute"
	Cathedra   Organization = "cathedra"
)

func (org Organization) Value() (driver.Value, error) {
	return string(org), nil
}
func (org *Organization) Scan(value interface{}) error {
	if v, ok := value.(string); ok {
		*org = Organization(v)
		return nil
	}
	return errors.New("Organization scan failed")
}
func (org Organization) String() string {
	return string(org)
}
