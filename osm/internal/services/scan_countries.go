package services

import (
	"encoding/json"
	"os"

	"github.com/paulmach/osm"
)

func ScanCountries(node *osm.Node, countries *os.File) error {
	if node.Tags.Find("place") == "country" {
		data, _ := json.Marshal(node)

		countries.WriteString(string(data))
		countries.WriteString("\n")
	}

	return nil
}
