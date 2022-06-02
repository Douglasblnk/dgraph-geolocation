package models

type Country struct {
	Name         string
	Latitude     float64
	Longitude    float64
	Capital_city *string
	Country_code *string
	Population   *string
}
