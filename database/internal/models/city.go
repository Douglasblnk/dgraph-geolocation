package models

type City struct {
	Name        string `json:"City.name"`
	Latitude    string `json:"City.latitude"`
	Longitude   string `json:"City.longitude"`
	CountryName string `json:"City.country_name"`
	CountryCode string `json:"City.country_code"`
	StateName   string `json:"City.state_name"`
	StateCode   string `json:"City.state_code"`
}
