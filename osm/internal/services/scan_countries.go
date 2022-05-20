package services

import (
	"github.com/paulmach/osm"
)

func ScanCountries(node *osm.Node) {

	if node.Tags.Find("place") == "continent" {

	}

}
