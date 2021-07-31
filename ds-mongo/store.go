package dsmongo

import (
	"context"
	"crypto/sha256"
	"fmt"

	dsrpc "github.com/beeleelee/go-ds-rpc"
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
			Msg: err.Error(),
		}
		if err == mongo.ErrNoDocuments {
			r.Code = dsrpc.ErrCode_ErrNotFound
		}
		return r, err
	}
	return &dsrpc.CommonReply{}, nil
}

func (ms *MongoStore) Delete(ctx context.Context, req *dsrpc.CommonRequest) (*dsrpc.CommonReply, error) {
	err := ms.client.Delete(ctx, req.GetKey())
	if err != nil {
		r := &dsrpc.CommonReply{
			Msg: err.Error(),
		}
		if err == mongo.ErrNoDocuments {
			r.Code = dsrpc.ErrCode_ErrNotFound
		}
		return r, err
	}
	return &dsrpc.CommonReply{}, nil
}

func (ms *MongoStore) Get(ctx context.Context, req *dsrpc.CommonRequest) (*dsrpc.CommonReply, error) {
	v, err := ms.client.Get(ctx, req.GetKey())
	if err != nil {
		// logging.Info("=====")
		// logging.Warn(err)
		// logging.Info("=====")
		r := &dsrpc.CommonReply{
			Msg: err.Error(),
		}
		if err == mongo.ErrNoDocuments {
			logging.Info("mongo not found")
			r.Code = dsrpc.ErrCode_ErrNotFound
		}
		logging.Infof("%v, %v", r.GetCode(), r.GetMsg())
		return r, err
	}
	return &dsrpc.CommonReply{Value: v}, nil
}

func (ms *MongoStore) Has(ctx context.Context, req *dsrpc.CommonRequest) (*dsrpc.CommonReply, error) {
	has, err := ms.client.Has(ctx, req.GetKey())
	if err != nil {
		r := &dsrpc.CommonReply{
			Msg: err.Error(),
		}
		if err == mongo.ErrNoDocuments {
			r.Code = dsrpc.ErrCode_ErrNotFound
		}
		return r, err
	}
	return &dsrpc.CommonReply{Success: has}, nil
}

func (ms *MongoStore) GetSize(ctx context.Context, req *dsrpc.CommonRequest) (*dsrpc.CommonReply, error) {
	v, err := ms.client.GetSize(ctx, req.GetKey())
	if err != nil {
		r := &dsrpc.CommonReply{
			Msg: err.Error(),
		}
		if err == mongo.ErrNoDocuments {
			r.Code = dsrpc.ErrCode_ErrNotFound
		}
		return r, err
	}
	return &dsrpc.CommonReply{Size: v}, nil
}

func (ms *MongoStore) Query(ctx context.Context, req *dsrpc.QueryRequest) (*dsrpc.QueryReply, error) {

	return &dsrpc.QueryReply{Res: []byte{}}, nil
}

func (ms *MongoStore) Close(context.Context) error {
	return ms.client.Close()
}

func sha256String(d []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(d))
}
