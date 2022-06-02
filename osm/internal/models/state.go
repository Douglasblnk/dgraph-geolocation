package models

type State struct {
	Name          string
	Latitude      float64
	Longitude     float64
	Is_in_country *string
	State_code    *string
	Population    *string
}
