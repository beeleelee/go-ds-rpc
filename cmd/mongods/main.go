package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	dsrpc "github.com/beeleelee/go-ds-rpc"
	dsmongo "github.com/beeleelee/go-ds-rpc/ds-mongo"
	log "github.com/ipfs/go-log/v2"
	"google.golang.org/grpc"
)

var logging = log.Logger("mongods")

var (
	listenPort uint
	dbUri      string
	dbName     string
	collName   string
)

func init() {
	log.SetLogLevel("*", "info")
}

func main() {
	logging.Info("### 启动中... ###")
	flag.UintVar(&listenPort, "port", 1516, "rpc listen port")
	flag.StringVar(&dbUri, "db-uri", "", "db connection address")
	flag.StringVar(&dbName, "db-name", "", "db name")
	flag.StringVar(&collName, "coll-name", "", "db collection")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ms, err := dsmongo.NewMongoStore(dsmongo.Options{
		Uri:      dbUri,
		DBName:   dbName,
		CollName: collName,
	})
	if err != nil {
		logging.Fatal(err)
	}
	defer ms.Close(ctx)

	// 启动 pinner rpc服务
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", listenPort))
	if err != nil {
		logging.Fatalf("failed to listen: %v", err)
	}

	// TODO 使用https证书建立安全通道
	rpcSrv := grpc.NewServer()
	dsrpc.RegisterKVStoreServer(rpcSrv, ms)
	go func() {
		if err := rpcSrv.Serve(lis); err != nil && err != http.ErrServerClosed {
			logging.Fatalf("rpc listen: %s\n", err)
		}
	}()
	logging.Infof("rpc listen port: %d", listenPort)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logging.Info("Shutdown Server ...")

	rpcSrv.GracefulStop()

	logging.Info("Server exiting")
}
