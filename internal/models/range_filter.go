package models

// Generyczny model zakresu
type RangeFilter struct {
	Largest  int `json:"largest" doc:"Największa wartość zakresu" format:"int32"`
	Smallest int `json:"smallest" doc:"Najmniejsza wartość zakresu" format:"int32"`
}
