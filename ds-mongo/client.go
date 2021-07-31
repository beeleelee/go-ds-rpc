package dsmongo

import (
	dsrpc "github.com/beeleelee/go-ds-rpc"
	"google.golang.org/grpc"
)

var _ dsrpc.KVStoreClient = (*MongoStoreClient)(nil)

type MongoStoreClient struct {
	conn *grpc.ClientConn
	dsrpc.KVStoreClient
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
		conn:          conn,
		KVStoreClient: dsrpc.NewKVStoreClient(conn),
	}, nil
}

func (ms *MongoStoreClient) Close() error {
	return ms.conn.Close()
}
