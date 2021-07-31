package dsmongo

import (
	dsrpc "github.com/beeleelee/go-ds-rpc"
	"google.golang.org/grpc"
)

type MongoStoreClient struct {
	conn   *grpc.ClientConn
	Client dsrpc.KVStoreClient
}

func NewMongoStoreClient(srv string) (*MongoStoreClient, error) {
	if srv == "" {
		logging.Fatal("mongostore rpc server address is missing")
	}
	conn, err := grpc.Dial(srv, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &MongoStoreClient{
		conn:   conn,
		Client: dsrpc.NewKVStoreClient(conn),
	}, nil
}

func (ms *MongoStoreClient) Close() error {
	return ms.conn.Close()
}
