package dsrpc

import (
	context "context"

	ds "github.com/ipfs/go-datastore"
	dsq "github.com/ipfs/go-datastore/query"
	"golang.org/x/xerrors"
)

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
	r, err := d.client.Put(d.ctx, &PutRequest{
		Key:   k.String(),
		Value: value,
	})
	if err != nil {
		return err
	}
	if r.GetErr() != "" {
		return xerrors.New(r.GetErr())
	}
	return nil
}

func (d *DataStore) Get(k ds.Key) ([]byte, error) {
	r, err := d.client.Get(d.ctx, &StoreKey{
		Key: k.String(),
	})
	if err != nil {
		return nil, err
	}
	return r.GetValue(), nil
}

func (d *DataStore) Has(k ds.Key) (bool, error) {
	r, err := d.client.Has(d.ctx, &StoreKey{
		Key: k.String(),
	})
	if err != nil {
		return false, err
	}
	return r.GetSuccess(), nil
}

func (d *DataStore) GetSize(k ds.Key) (int, error) {
	r, err := d.client.GetSize(d.ctx, &StoreKey{
		Key: k.String(),
	})
	if err != nil {
		return 0, err
	}
	return int(r.GetSize()), nil
}

func (d *DataStore) Delete(k ds.Key) error {
	r, err := d.client.Delete(d.ctx, &StoreKey{
		Key: k.String(),
	})
	if err != nil {
		return err
	}
	if r.GetErr() != "" {
		return xerrors.New(r.GetErr())
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
