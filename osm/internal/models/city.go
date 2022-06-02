package models

import "github.com/paulmach/osm"

type City struct {
	ID              osm.NodeID
	Name            string
	Latitude        float64
	Longitude       float64
	Is_capital      string
	Is_in           *string
	Is_in_continent *string
	Is_in_country   *string
	Is_in_state     *string
	Population      *string
}
