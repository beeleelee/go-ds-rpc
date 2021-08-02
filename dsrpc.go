package dsrpc

import (
	context "context"

	ds "github.com/ipfs/go-datastore"
	dsq "github.com/ipfs/go-datastore/query"
	log "github.com/ipfs/go-log/v2"
	"golang.org/x/xerrors"
)

var logging = log.Logger("dsrpc")
var _ ds.Batching = (*DataStore)(nil)

type DataStore struct {
	ctx    context.Context
	client KVStoreClient
}

func NewDataStore(client KVStoreClient) (*DataStore, error) {
	if client == nil {
		return nil, xerrors.New("missing KVStoreClient instance")
	}
	return &DataStore{
		client: client,
		ctx:    context.Background(),
	}, nil
}

func (d *DataStore) Put(k ds.Key, value []byte) error {
	logging.Infof("p put key: %v", k)
	r, err := d.client.Put(d.ctx, &CommonRequest{
		Key:   k.String(),
		Value: value,
	})
	if err != nil {
		return err
	}
	if r.GetCode() != ErrCode_None {
		logging.Warnf("p put err code: %d, msg: %s", r.GetCode(), r.GetMsg())
		return xerrors.New(r.GetMsg())
	}
	return nil
}

func (d *DataStore) Get(k ds.Key) ([]byte, error) {
	logging.Infof("p get key: %v", k)
	r, err := d.client.Get(d.ctx, &CommonRequest{
		Key: k.String(),
	})
	if err != nil {
		return nil, err
	}
	if r.GetCode() != ErrCode_None {
		if r.GetCode() == ErrCode_ErrNotFound {
			return nil, ds.ErrNotFound
		}
		return nil, xerrors.New(r.GetMsg())
	}
	return r.GetValue(), nil
}

func (d *DataStore) Has(k ds.Key) (bool, error) {
	logging.Infof("p has key: %v", k)
	r, err := d.client.Has(d.ctx, &CommonRequest{
		Key: k.String(),
	})
	if err != nil {
		return false, err
	}
	if r.GetCode() != ErrCode_None {
		if r.GetCode() == ErrCode_ErrNotFound {
			return false, ds.ErrNotFound
		}
		return false, xerrors.New(r.GetMsg())
	}
	return r.GetSuccess(), nil
}

func (d *DataStore) GetSize(k ds.Key) (int, error) {
	logging.Infof("p getsize key: %v", k)
	r, err := d.client.GetSize(d.ctx, &CommonRequest{
		Key: k.String(),
	})
	if err != nil {
		return 0, err
	}
	if r.GetCode() != ErrCode_None {
		if r.GetCode() == ErrCode_ErrNotFound {
			return 0, ds.ErrNotFound
		}
		return 0, xerrors.New(r.GetMsg())
	}
	return int(r.GetSize()), nil
}

func (d *DataStore) Delete(k ds.Key) error {
	logging.Infof("p delete key: %v", k)
	r, err := d.client.Delete(d.ctx, &CommonRequest{
		Key: k.String(),
	})
	if err != nil {
		return err
	}
	if r.GetCode() != ErrCode_None {
		if r.GetCode() == ErrCode_ErrNotFound {
			return ds.ErrNotFound
		}
		return xerrors.New(r.GetMsg())
	}

	return nil
}

func (d *DataStore) Sync(ds.Key) error {
	return nil
}

func (d *DataStore) Close() error {
	return nil
}

func (d *DataStore) Query(q dsq.Query) (dsq.Results, error) {
	return nil, nil
}

func (d *DataStore) Batch() (ds.Batch, error) {
	return ds.NewBasicBatch(d), nil
}
