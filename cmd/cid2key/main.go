package main

import (
	"flag"
	"fmt"

	gcid "github.com/ipfs/go-cid"
	ds "github.com/ipfs/go-datastore"
	dshelp "github.com/ipfs/go-ipfs-ds-help"
	log "github.com/ipfs/go-log/v2"
)

var logging = log.Logger("mongods")

var (
	cid string
	key string
)

func init() {
	log.SetLogLevel("*", "info")
}

func main() {
	flag.StringVar(&cid, "cid", "", "")
	flag.StringVar(&key, "key", "", "datastore key")
	flag.Parse()

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	if cid == "" && key == "" {
		fmt.Println("--cid or --key should be provided")
		return
	}
	if cid != "" {
		id, err := gcid.Decode(cid)
		if err != nil {
			logging.Fatal(err)
			return
		}
		logging.Infof("key: %s", dshelp.MultihashToDsKey(id.Hash()))
		return
	}
	if key != "" {
		id, err := dshelp.DsKeyToCidV1(ds.NewKey(key), gcid.Raw)
		if err != nil {
			logging.Fatal(err)
			return
		}
		logging.Infof("cid: %s", id)
		return
	}
}
