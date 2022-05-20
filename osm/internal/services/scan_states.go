package services

import (
	"encoding/json"
	"os"

	"github.com/paulmach/osm"
)

func ScanStates(node *osm.Node, states *os.File) error {
	if node.Tags.Find("place") == "state" {
		data, _ := json.Marshal(node)

		states.WriteString(string(data))
		states.WriteString("\n")
	}

	return nil
}
