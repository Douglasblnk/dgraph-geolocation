package services

import (
	"dgraph-osm/osm/internal/models"
	"dgraph-osm/osm/internal/utils"
	"encoding/json"
	"os"

	"github.com/paulmach/osm"
)

func ScanStates(node *osm.Node, states *os.File) *models.State {
	var stateModel *models.State

	if node.Tags.Find("place") == "state" {
		stateModel = &models.State{
			Name:      node.Tags.Find("name"),
			Latitude:  node.Lat,
			Longitude: node.Lon,
			Is_in_country: utils.StringPointer(
				utils.Ternary(
					node.Tags.Find("is_in:country"),
					node.Tags.Find("is_in"),
				),
			),
			State_code: utils.StringPointer(
				utils.Ternary(
					node.Tags.Find("state_code"),
					node.Tags.Find("ref"),
				),
			),
			Population: utils.StringPointer(node.Tags.Find("population")),
		}

		data, _ := json.Marshal(stateModel)

		states.WriteString(string(data))
		states.WriteString("\n")
	}

	return stateModel
}
