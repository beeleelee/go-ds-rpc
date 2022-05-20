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

type DataStore struct {
	client KVStoreClient
}

var _ds DataStore
var _ ds.Batching = _ds

func NewDataStore(client KVStoreClient) (*DataStore, error) {
	if client == nil {
		return nil, xerrors.New("missing KVStoreClient instance")
	}
	return &DataStore{
		client: client,
	}, nil
}

func (d DataStore) Put(ctx context.Context, k ds.Key, value []byte) error {
	r, err := d.client.Put(ctx, &CommonRequest{
		Key:   k.String(),
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

func (d DataStore) Get(ctx context.Context, k ds.Key) ([]byte, error) {
	r, err := d.client.Get(ctx, &CommonRequest{
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

func (d DataStore) Has(ctx context.Context, k ds.Key) (bool, error) {
	r, err := d.client.Has(ctx, &CommonRequest{
		Key: k.String(),
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

func (d DataStore) GetSize(ctx context.Context, k ds.Key) (int, error) {
	r, err := d.client.GetSize(ctx, &CommonRequest{
		Key: k.String(),
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

func (d DataStore) Delete(ctx context.Context, k ds.Key) error {
	r, err := d.client.Delete(ctx, &CommonRequest{
		Key: k.String(),
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

func (d DataStore) Sync(ctx context.Context, k ds.Key) error {
	return nil
}

func (d DataStore) Close() error {
	return nil
}

func (d DataStore) Query(ctx context.Context, q dsq.Query) (dsq.Results, error) {
	b, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}
	r, err := d.client.Query(ctx, &QueryRequest{
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

func (d DataStore) Batch(ctx context.Context) (ds.Batch, error) {
	return ds.NewBasicBatch(d), nil
}
