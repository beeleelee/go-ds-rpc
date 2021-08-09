package dsmongo

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	dsrpc "github.com/beeleelee/go-ds-rpc"
	dsq "github.com/ipfs/go-datastore/query"
	"go.mongodb.org/mongo-driver/mongo"
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

func (ms *MongoStore) Put(ctx context.Context, req *dsrpc.CommonRequest) (*dsrpc.CommonReply, error) {
	hk := sha256String(req.GetValue())

	refItem := &RefItem{
		ID:  req.GetKey(),
		Ref: hk,
	}
	storeItem := &StoreItem{
		ID:    hk,
		Value: req.GetValue(),
	}
	err := ms.client.Put(ctx, storeItem, refItem)
	if err != nil {
		r := &dsrpc.CommonReply{
			Msg:  err.Error(),
			Code: dsrpc.ErrCode_Others,
		}
		if err == mongo.ErrNoDocuments {
			r.Code = dsrpc.ErrCode_ErrNotFound
		}
		return r, nil
	}
	return &dsrpc.CommonReply{}, nil
}

func (ms *MongoStore) Delete(ctx context.Context, req *dsrpc.CommonRequest) (*dsrpc.CommonReply, error) {
	err := ms.client.Delete(ctx, req.GetKey())
	if err != nil {
		r := &dsrpc.CommonReply{
			Msg:  err.Error(),
			Code: dsrpc.ErrCode_Others,
		}
		if err == mongo.ErrNoDocuments {
			r.Code = dsrpc.ErrCode_ErrNotFound
		}
		return r, nil
	}
	return &dsrpc.CommonReply{}, nil
}

func (ms *MongoStore) Get(ctx context.Context, req *dsrpc.CommonRequest) (*dsrpc.CommonReply, error) {
	v, err := ms.client.Get(ctx, req.GetKey())
	if err != nil {
		r := &dsrpc.CommonReply{
			Msg:  err.Error(),
			Code: dsrpc.ErrCode_Others,
		}
		if err == mongo.ErrNoDocuments {
			r.Code = dsrpc.ErrCode_ErrNotFound
		}
		return r, nil
	}
	return &dsrpc.CommonReply{Value: v}, nil
}

func (ms *MongoStore) Has(ctx context.Context, req *dsrpc.CommonRequest) (*dsrpc.CommonReply, error) {
	has, err := ms.client.Has(ctx, req.GetKey())
	if err != nil {
		r := &dsrpc.CommonReply{
			Msg:  err.Error(),
			Code: dsrpc.ErrCode_Others,
		}
		if err == mongo.ErrNoDocuments {
			r.Code = dsrpc.ErrCode_ErrNotFound
		}
		return r, nil
	}
	return &dsrpc.CommonReply{Success: has}, nil
}

func (ms *MongoStore) GetSize(ctx context.Context, req *dsrpc.CommonRequest) (*dsrpc.CommonReply, error) {
	v, err := ms.client.GetSize(ctx, req.GetKey())
	if err != nil {
		r := &dsrpc.CommonReply{
			Msg:  err.Error(),
			Code: dsrpc.ErrCode_Others,
		}
		if err == mongo.ErrNoDocuments {
			r.Code = dsrpc.ErrCode_ErrNotFound
		}
		//logging.Infof("get size error, code: %v, msg: %v", r.GetCode(), r.GetMsg())
		return r, nil
	}

	return &dsrpc.CommonReply{Size: v}, nil
}

func (ms *MongoStore) Query(req *dsrpc.QueryRequest, reply dsrpc.KVStore_QueryServer) error {
	re := dsq.Query{}
	err := json.Unmarshal(req.GetQ(), &re)
	if err != nil {
		return err
	}
	logging.Infof("query: %s", re)
	items, err := ms.client.Query(context.Background(), re)
	if err != nil {
		return err
	}
	for ent := range items {
		b, err := json.Marshal(ent)
		if err != nil {
			return err
		}
		err = reply.Send(&dsrpc.QueryReply{
			Res: b,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (ms *MongoStore) Close(context.Context) error {
	return ms.client.Close()
}

func sha256String(d []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(d))
}
