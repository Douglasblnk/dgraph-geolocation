package services

import (
	"context"
	"dgraph-osm/database/internal/connection"
	"dgraph-osm/database/internal/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/dgraph-io/dgo/v210/protos/api"
)

type States struct {
	Name        string  `json:"name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Country     string  `json:"country"`
	CountryCode string  `json:"country_code"`
	StateCode   string  `json:"state_code"`
}

func getStates() []*models.State {
	stateFile, _ := os.Open("/home/douglasblnk/Documentos/Projects/dgraph-go/database/data/states.json")
	defer stateFile.Close()

	statesByteValue, _ := ioutil.ReadAll(stateFile)

	stateModel := []*States{}
	json.Unmarshal(statesByteValue, &stateModel)

	var states []*models.State

	for _, state := range stateModel {
		s := &models.State{
			Name:        state.Name,
			Latitude:    state.Latitude,
			Longitude:   state.Longitude,
			Country:     state.Country,
			CountryCode: state.CountryCode,
			StateCode:   state.StateCode,
		}

		states = append(states, s)
	}

	return states
}

func InsertStates() {
	dgcon, close := connection.DgraphConnection()
	defer close()

	states := getStates()

	json, _ := json.Marshal(&states)

	mu := &api.Mutation{
		CommitNow: true,
		SetJson:   json,
	}

	response, err := dgcon.NewTxn().Mutate(context.Background(), mu)

	if err != nil {
		log.Fatal("err", err)
	}

	println(response)
}
