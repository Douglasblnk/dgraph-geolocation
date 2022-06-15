package services

import (
	"context"
	"dgraph-osm/osm/internal/models"
	"dgraph-osm/osm/internal/utils"
	"encoding/json"
	"log"
	"os"

	"github.com/paulmach/osm"
	"github.com/paulmach/osm/osmapi"
	"github.com/paulmach/osm/osmgeojson"
)

func ScanStates(node *osm.Node, states *os.File) *models.State {
	var stateModel *models.State
	delta := 0.0001
	ctx := context.Background()

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

		// data, _ := json.Marshal(&stateModel)

		bounds := &osm.Bounds{
			MinLat: stateModel.Latitude - delta,
			MaxLat: stateModel.Latitude + delta,
			MinLon: stateModel.Longitude - delta,
			MaxLon: stateModel.Longitude + delta,
		}

		o, _ := osmapi.Map(ctx, bounds)

		fc, err := osmgeojson.Convert(o)

		if err != nil {
			log.Fatal("err =>", err)
		}

		gj, _ := json.MarshalIndent(fc, "", " ")

		states.WriteString(string(gj))
	}

	return stateModel
}
