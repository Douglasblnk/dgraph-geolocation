package services

import "os"

type Files struct {
	Country *os.File
	State   *os.File
	City    *os.File
}

func CreateFile() (*Files, error) {
	var err error
	files := &Files{}

	files.Country, err = os.Create("osm/data/json/countries.json")

	if err != nil {
		return nil, err
	}

	files.State, err = os.Create("osm/data/json/states.json")

	if err != nil {
		return nil, err
	}

	files.City, err = os.Create("osm/data/json/cities.json")

	if err != nil {
		return nil, err
	}

	return files, nil
}
