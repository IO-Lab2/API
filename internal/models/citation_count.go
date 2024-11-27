package models

type CitationsFilter struct {
	Largest  int `json:"largest"`
	Smallest int `json:"smallest"`
}
