package main

import (
	"flag"
	"fmt"

	"finishy1995/device-manager/enhanced/processor/internal/config"
	"finishy1995/device-manager/enhanced/processor/internal/server"
	"finishy1995/device-manager/enhanced/processor/internal/svc"
	"finishy1995/device-manager/enhanced/processor/pb/processor"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/processor.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		processor.RegisterProcessorServer(grpcServer, server.NewProcessorServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
