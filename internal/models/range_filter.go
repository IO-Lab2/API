package models

// Generyczny model zakresu
type RangeFilter struct {
	Largest  int `json:"largest"`
	Smallest int `json:"smallest"`
}
