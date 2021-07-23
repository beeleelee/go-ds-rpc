package dsmongo_test

import (
	"context"
	"testing"

	"github.com/beeleelee/dsrpc"
	dsmongo "github.com/beeleelee/dsrpc/ds-mongo"
)

func TestMongoStore(t *testing.T) {
	rpc_uri := "127.0.0.1:1516"
	client, err := dsmongo.NewMongoStoreClient(rpc_uri)
	if err != nil {
		t.Fatal(err)
	}
	dlen := 3 << 20
	d := make([]byte, dlen)

	ctx := context.Background()
	r, err := client.Client.Put(ctx, &dsrpc.PutRequest{
		Key:   "test-data-1GiB",
		Value: d,
	})
	if err != nil {
		t.Fatal(err)
	}
	if r.GetErr() != "" {
		t.Fatal(r.GetErr())
	}
	sr, err := client.Client.GetSize(ctx, &dsrpc.StoreKey{
		Key: "test-data-1GiB",
	})

	if err != nil {
		t.Fatal(err)
	}
	if sr.GetSize() != 3<<20 {
		t.Fatal("data size not match")
	}
}
