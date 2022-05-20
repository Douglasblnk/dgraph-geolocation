package main

import (
	"dgraph-osm/osm/internal/connection"
	"dgraph-osm/osm/internal/models"
	"log"

	"github.com/dolan-in/dgman/v2"
)

func main() {
	dg, cancel := connection.DgraphConnection()
	defer cancel()

	model := &models.Continent{}

	schema, err := dgman.MutateSchema(dg, model)

	if err != nil {
		log.Fatal("err =>", err)
	}

	log.Printf("Schemas criados/atualizados com sucesso: \r\n %v \r\n", schema)

}
