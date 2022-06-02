package connection

import (
	"log"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"
)

func DgraphConnection() (*dgo.Dgraph, func() error) {
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure(), grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(32*10e6),
		grpc.MaxCallSendMsgSize(32*10e6),
	))

	if err != nil {
		log.Fatal("err =>", err)
	}

	return dgo.NewDgraphClient(api.NewDgraphClient(conn)), conn.Close
}
