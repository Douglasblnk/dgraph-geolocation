package services

import (
	"encoding/json"
	"os"

	"github.com/paulmach/osm"
)

var cityTypes []string = []string{
	"city",
	"district",
	"county",
	"municipality",
	"town",
	"village",
}

func ScanCities(node *osm.Node, cities *os.File) error {
	for _, cityType := range cityTypes {
		if node.Tags.Find("place") == cityType {
			data, _ := json.Marshal(node)

			cities.WriteString(string(data))
			cities.WriteString("\n")
		}
	}

	return nil
}
