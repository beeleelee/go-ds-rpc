package dsrpc_test

import (
	"context"
	"testing"
	"time"

	"crypto/rand"

	dsrpc "github.com/beeleelee/go-ds-rpc"
	dsmongo "github.com/beeleelee/go-ds-rpc/ds-mongo"
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
