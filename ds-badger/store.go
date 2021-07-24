package dsbadger

import (
	"context"

	"github.com/beeleelee/dsrpc"
	ds "github.com/ipfs/go-datastore"
	badgerds3 "github.com/textileio/go-ds-badger3"
)

type BadgerStore struct {
	dsrpc.UnimplementedKVStoreServer
	db *badgerds3.Datastore
}

func NewBadgerStore(path string) (*BadgerStore, error) {
	db, err := badgerds3.NewDatastore(path, nil)
	if err != nil {
		return nil, err
	}
	return &BadgerStore{
		db: db,
	}, nil
}

func (ms *BadgerStore) Put(ctx context.Context, req *dsrpc.PutRequest) (*dsrpc.ErrReply, error) {

	err := ms.db.Put(ds.NewKey(req.GetKey()), req.GetValue())

	if err != nil {
		return &dsrpc.ErrReply{
			Err: err.Error(),
		}, err
	}
	return &dsrpc.ErrReply{}, nil
}

func (ms *BadgerStore) Delete(ctx context.Context, req *dsrpc.StoreKey) (*dsrpc.ErrReply, error) {
	err := ms.db.Delete(ds.NewKey(req.GetKey()))
	if err != nil {
		return &dsrpc.ErrReply{
			Err: err.Error(),
		}, err
	}
	return &dsrpc.ErrReply{}, nil
}

func (ms *BadgerStore) Get(ctx context.Context, req *dsrpc.StoreKey) (*dsrpc.StoreValue, error) {
	v, err := ms.db.Get(ds.NewKey(req.GetKey()))
	if err != nil {
		return nil, err
	}
	return &dsrpc.StoreValue{Value: v}, nil
}

func (ms *BadgerStore) Has(ctx context.Context, req *dsrpc.StoreKey) (*dsrpc.BoolReply, error) {
	has, err := ms.db.Has(ds.NewKey(req.GetKey()))
	if err != nil {
		return &dsrpc.BoolReply{Success: false}, err
	}
	return &dsrpc.BoolReply{Success: has}, nil
}

func (ms *BadgerStore) GetSize(ctx context.Context, req *dsrpc.StoreKey) (*dsrpc.SizeReply, error) {
	v, err := ms.db.GetSize(ds.NewKey(req.GetKey()))
	if err != nil {
		return nil, err
	}
	return &dsrpc.SizeReply{Size: int64(v)}, nil
}

func (ms *BadgerStore) Query(ctx context.Context, req *dsrpc.StoreQuery) (*dsrpc.QueryReply, error) {
	return &dsrpc.QueryReply{Res: []byte{}}, nil
}

func (ms *BadgerStore) Close(context.Context) error {
	return ms.db.Close()
}
