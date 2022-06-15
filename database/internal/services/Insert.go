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

func LoadFiles() ([]*Country, []*States, []*Cities) {
	countryFile, _ := os.Open("/home/douglasblnk/Documentos/Projects/dgraph-go/database/data/countries.json")
	stateFile, _ := os.Open("/home/douglasblnk/Documentos/Projects/dgraph-go/database/data/states.json")
	cityFile, _ := os.Open("/home/douglasblnk/Documentos/Projects/dgraph-go/database/data/cities.json")

	defer countryFile.Close()
	defer stateFile.Close()
	defer cityFile.Close()

	countriesByteValue, _ := ioutil.ReadAll(countryFile)
	statesByteValue, _ := ioutil.ReadAll(stateFile)
	citiesByteValue, _ := ioutil.ReadAll(cityFile)

	countryModel := []*Country{}
	stateModel := []*States{}
	cityModel := []*Cities{}

	json.Unmarshal(countriesByteValue, &countryModel)
	json.Unmarshal(statesByteValue, &stateModel)
	json.Unmarshal(citiesByteValue, &cityModel)

	return countryModel, stateModel, cityModel
}

func GetStates(states []*States, countryName string) []*models.State {
	var s []*models.State

	for _, state := range states {
		if state.CountryName == countryName {
			model := &models.State{
				Name:        state.Name,
				Latitude:    state.Latitude,
				Longitude:   state.Longitude,
				CountryName: state.CountryName,
				CountryCode: state.CountryCode,
				StateCode:   state.StateCode,
			}

			s = append(s, model)
		}
	}

	return s
}

func GetCities(cities []*Cities, stateName string) []*models.City {
	var c []*models.City

	for _, city := range cities {
		if city.StateName == stateName {
			model := &models.City{
				Name:        city.Name,
				Latitude:    city.Latitude,
				Longitude:   city.Longitude,
				CountryName: city.CountryName,
				CountryCode: city.CountryCode,
				StateName:   city.StateName,
				StateCode:   city.StateCode,
			}

			c = append(c, model)
		}
	}

	return c
}

func Insert() {
	dgcon, close := connection.DgraphConnection()
	defer close()

	countries, states, cities := LoadFiles()

	for _, country := range countries {
		if country.Iso2 == "BR" {
			countryModel := &models.Country{
				Name:           country.Name,
				Capital:        country.Capital,
				Iso2:           country.Iso2,
				Iso3:           country.Iso3,
				Currency:       country.Currency,
				CurrencySymbol: country.CurrencySymbol,
				Emoji:          country.Emoji,
			}

			statesFromCountry := GetStates(states, country.Name)

			for index, state := range statesFromCountry {
				citiesFromState := GetCities(cities, state.Name)

				statesFromCountry[index].Cities = citiesFromState
			}

			countryModel.States = statesFromCountry

			json, _ := json.Marshal(countryModel)

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
	}
}
