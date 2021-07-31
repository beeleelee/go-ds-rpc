package dsmongo

import (
	"context"

	dsrpc "github.com/beeleelee/go-ds-rpc"
)

type MongoStore struct {
	dsrpc.UnimplementedKVStoreServer
	client *DSMongo
}

func NewMongoStore(opts Options) (*MongoStore, error) {
	cl, err := NewDSMongo(opts)
	if err != nil {
		return nil, err
	}
	return &MongoStore{
		client: cl,
	}, nil
}

func (ms *MongoStore) Put(ctx context.Context, req *dsrpc.PutRequest) (*dsrpc.ErrReply, error) {
	storeItem := &StoreItem{
		ID:    req.GetKey(),
		Value: req.GetValue(),
		Size:  int64(len(req.GetValue())),
	}
	err := ms.client.Put(ctx, storeItem)
	if err != nil {
		return &dsrpc.ErrReply{
			Err: err.Error(),
		}, err
	}
	return &dsrpc.ErrReply{}, nil
}

func (ms *MongoStore) Delete(ctx context.Context, req *dsrpc.StoreKey) (*dsrpc.ErrReply, error) {

	err := ms.client.Delete(ctx, req.GetKey())
	if err != nil {
		return &dsrpc.ErrReply{
			Err: err.Error(),
		}, err
	}
	return &dsrpc.ErrReply{}, nil
}

func (ms *MongoStore) Get(ctx context.Context, req *dsrpc.StoreKey) (*dsrpc.StoreValue, error) {

	v, err := ms.client.Get(ctx, req.GetKey())
	if err != nil {
		return nil, err
	}
	return &dsrpc.StoreValue{Value: v}, nil
}

func (ms *MongoStore) Has(ctx context.Context, req *dsrpc.StoreKey) (*dsrpc.BoolReply, error) {

	has, err := ms.client.Has(ctx, req.GetKey())
	if err != nil {
		return &dsrpc.BoolReply{Success: false}, err
	}
	return &dsrpc.BoolReply{Success: has}, nil
}

func (ms *MongoStore) GetSize(ctx context.Context, req *dsrpc.StoreKey) (*dsrpc.SizeReply, error) {

	v, err := ms.client.GetSize(ctx, req.GetKey())
	if err != nil {
		return nil, err
	}
	return &dsrpc.SizeReply{Size: v}, nil
}

func (ms *MongoStore) Query(ctx context.Context, req *dsrpc.StoreQuery) (*dsrpc.QueryReply, error) {

	return &dsrpc.QueryReply{Res: []byte{}}, nil
}

func (ms *MongoStore) Close(context.Context) error {
	return ms.client.Close()
}
