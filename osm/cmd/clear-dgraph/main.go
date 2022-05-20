package main

import (
	"context"
	"dgraph-osm/osm/internal/connection"
	"log"

	"github.com/dgraph-io/dgo/v210/protos/api"
)

func main() {
	dg, cancel := connection.DgraphConnection()
	defer cancel()

	ctx := context.Background()

	if err := dg.Alter(ctx, &api.Operation{DropAll: true}); err != nil {
		log.Fatal("Error ao limpar banco de dados", err)
	}

	log.Println("Banco de dados limpo.")
}
