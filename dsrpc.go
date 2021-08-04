package dsrpc

import (
	context "context"
	"encoding/json"

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
	prefix ds.Key
}

func NewDataStore(prefix string, client KVStoreClient) (*DataStore, error) {
	if client == nil {
		return nil, xerrors.New("missing KVStoreClient instance")
	}
	return &DataStore{
		client: client,
		ctx:    context.Background(),
		prefix: ds.NewKey(prefix),
	}, nil
}

func (d *DataStore) Put(k ds.Key, value []byte) error {
	r, err := d.client.Put(d.ctx, &CommonRequest{
		Key:   d.prefix.String() + k.String(),
		Value: value,
	})
	if err != nil {
		return err
	}
	if r.GetCode() != ErrCode_None {
		return xerrors.New(r.GetMsg())
	}
	return nil
}

func (d *DataStore) Get(k ds.Key) ([]byte, error) {
	r, err := d.client.Get(d.ctx, &CommonRequest{
		Key: d.prefix.String() + k.String(),
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
	r, err := d.client.Has(d.ctx, &CommonRequest{
		Key: d.prefix.String() + k.String(),
	})
	if err != nil {
		return false, err
	}
	if r.GetCode() != ErrCode_None {
		if r.GetCode() == ErrCode_ErrNotFound {
			return false, nil
		}
		return false, xerrors.New(r.GetMsg())
	}
	return r.GetSuccess(), nil
}

func (d *DataStore) GetSize(k ds.Key) (int, error) {
	r, err := d.client.GetSize(d.ctx, &CommonRequest{
		Key: d.prefix.String() + k.String(),
	})
	if err != nil {
		return -1, err
	}
	if r.GetCode() != ErrCode_None {
		if r.GetCode() == ErrCode_ErrNotFound {
			return -1, ds.ErrNotFound
		}
		return -1, xerrors.New(r.GetMsg())
	}
	return int(r.GetSize()), nil
}

func (d *DataStore) Delete(k ds.Key) error {
	r, err := d.client.Delete(d.ctx, &CommonRequest{
		Key: d.prefix.String() + k.String(),
	})
	if err != nil {
		return err
	}
	if r.GetCode() != ErrCode_None {
		if r.GetCode() == ErrCode_ErrNotFound {
			//return ds.ErrNotFound
			return nil
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
	q.Prefix = d.prefix.String() + q.Prefix
	b, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}
	r, err := d.client.Query(d.ctx, &QueryRequest{
		Q: b,
	})
	if err != nil {
		r.CloseSend()
		return nil, err
	}

	nextValue := func() (dsq.Result, bool) {
		ritem, err := r.Recv()
		// if err == io.EOF {
		// 	return dsq.Result{}, false
		// }
		if err != nil {
			return dsq.Result{Error: err}, false
		}

		ent := dsq.Entry{}
		err = json.Unmarshal(ritem.GetRes(), &ent)
		if err != nil {
			return dsq.Result{Error: err}, false
		}

		return dsq.Result{Entry: ent}, true
	}

	return dsq.ResultsFromIterator(q, dsq.Iterator{
		Close: func() error {
			return r.CloseSend()
		},
		Next: nextValue,
	}), nil
}

func (d *DataStore) Batch() (ds.Batch, error) {
	return ds.NewBasicBatch(d), nil
}
