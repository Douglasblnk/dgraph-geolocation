package connection

import (
	"log"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"
)

func DgraphConnection() (*dgo.Dgraph, func() error) {
	conn, err := grpc.Dial("127.0.0.1:9080", grpc.WithInsecure())

	if err != nil {
		log.Fatal("err =>", err)
	}

	return dgo.NewDgraphClient(api.NewDgraphClient(conn)), conn.Close
}
