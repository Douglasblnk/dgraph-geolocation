package models

type City struct {
	Name        string  `json:"City.Name"`
	Latitude    float64 `json:"City.Latitude"`
	Longitude   float64 `json:"City.Longitude"`
	Country     string  `json:"City.Country_name"`
	CountryCode string  `json:"City.Country_code"`
	State       string  `json:"City.State_name"`
	StateCode   string  `json:"City.State_code"`
}
