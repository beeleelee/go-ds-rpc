package dsrpc_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"crypto/rand"

	dsrpc "github.com/beeleelee/go-ds-rpc"
	dsmongo "github.com/beeleelee/go-ds-rpc/ds-mongo"
	dsq "github.com/ipfs/go-datastore/query"
	dag "github.com/ipfs/go-merkledag"
)

func TestMongoStore(t *testing.T) {
	rpc_uri := "127.0.0.1:1520"
	client, err := dsmongo.NewMongoStoreClient(rpc_uri)
	if err != nil {
		t.Fatal(err)
	}
	dataNum := 1500
	dataSize := 1 << 20

	ctx := context.Background()
	dagList := make([]*dag.ProtoNode, dataNum)

	for i, _ := range dagList {
		d := make([]byte, dataSize)
		rand.Read(d)
		dagList[i] = dag.NodeWithData(d)
	}
	putStart := time.Now()
	for _, dn := range dagList {
		_, err := client.Put(ctx, &dsrpc.CommonRequest{
			Key:   dn.Cid().String(),
			Value: dn.Data(),
		})
		if err != nil {
			t.Fatal(err)
		}
	}
	t.Logf("put time elapsed: %v", time.Now().Sub(putStart))

	getStart := time.Now()
	for _, dn := range dagList {
		_, err := client.Get(ctx, &dsrpc.CommonRequest{
			Key: dn.Cid().String(),
		})
		if err != nil {
			t.Fatal(err)
		}
	}
	t.Logf("get time elapsed: %v", time.Now().Sub(getStart))

	getSizeStart := time.Now()
	for _, dn := range dagList {
		_, err := client.GetSize(ctx, &dsrpc.CommonRequest{
			Key: dn.Cid().String(),
		})
		if err != nil {
			t.Fatal(err)
		}
	}
	t.Logf("get size time elapsed: %v", time.Now().Sub(getSizeStart))

	hasStart := time.Now()
	for _, dn := range dagList {
		_, err := client.Has(ctx, &dsrpc.CommonRequest{
			Key: dn.Cid().String(),
		})
		if err != nil {
			t.Fatal(err)
		}
	}
	t.Logf("has time elapsed: %v", time.Now().Sub(hasStart))

	deleteStart := time.Now()
	for _, dn := range dagList {
		_, err := client.Delete(ctx, &dsrpc.CommonRequest{
			Key: dn.Cid().String(),
		})
		if err != nil {
			t.Fatal(err)
		}
	}
	t.Logf("delete time elapsed: %v", time.Now().Sub(deleteStart))
}

// func TestGet(t *testing.T) {
// 	rpc_uri := "127.0.0.1:1520"
// 	client, err := dsmongo.NewMongoStoreClient(rpc_uri)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// r, err := client.Get(context.Background(), &dsrpc.CommonRequest{
// 	// 	Key: "/pins/state/dirty",
// 	// })
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	// if r.GetCode() != dsrpc.ErrCode_None {
// 	// 	t.Fatal(r.GetMsg())
// 	// }

// 	_, err = client.Put(context.Background(), &dsrpc.CommonRequest{
// 		Key:   "/local/filesroot",
// 		Value: []byte("test data"),
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

func TestQuery(t *testing.T) {
	rpc_uri := "127.0.0.1:1520"
	client, err := dsmongo.NewMongoStoreClient(rpc_uri)
	if err != nil {
		t.Fatal(err)
	}
	q := &dsq.Query{
		Prefix: ".*",
	}
	qb, err := json.Marshal(q)
	if err != nil {
		t.Fatal(err)
	}
	r, err := client.Query(context.TODO(), &dsrpc.QueryRequest{
		Q: qb,
	})
	if err != nil {
		t.Fatal(err)
	}
	for {
		qre, err := r.Recv()
		if err != nil {
			t.Fatal(err)
		}
		ent := dsq.Entry{}
		err = json.Unmarshal(qre.GetRes(), &ent)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(ent.Key, ent.Size)
	}

}
