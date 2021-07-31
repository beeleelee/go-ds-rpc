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
	dsbadager "github.com/beeleelee/go-ds-rpc/ds-badger"
	log "github.com/ipfs/go-log/v2"
	"google.golang.org/grpc"
)

var logging = log.Logger("badgerds")

var (
	listenPort uint
	dbPath     string
)

func init() {
	log.SetLogLevel("*", "info")
}

func main() {
	logging.Info("### 启动中... ###")
	flag.UintVar(&listenPort, "port", 1517, "rpc listen port")
	flag.StringVar(&dbPath, "path", "", "badger db path")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if dbPath == "" {
		logging.Fatal("badger db path is need")
	}
	logging.Infof("db path: %s", dbPath)
	bs, err := dsbadager.NewBadgerStore(dbPath)
	if err != nil {
		logging.Fatal(err)
	}
	defer bs.Close(ctx)

	// 启动 pinner rpc服务
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", listenPort))
	if err != nil {
		logging.Fatalf("failed to listen: %v", err)
	}

	// TODO 使用https证书建立安全通道
	rpcSrv := grpc.NewServer()
	dsrpc.RegisterKVStoreServer(rpcSrv, bs)
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
