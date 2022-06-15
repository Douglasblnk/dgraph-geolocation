package models

type State struct {
	Name        string  `json:"State.name"`
	Latitude    string  `json:"State.latitude"`
	Longitude   string  `json:"State.longitude"`
	CountryName string  `json:"State.country_name"`
	CountryCode string  `json:"State.country_code"`
	StateCode   string  `json:"State.state_code"`
	Cities      []*City `json:"State.cities"`
}
