package main

import (
	"flag"
	"fmt"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/email/internal/config"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/email/internal/server"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/email/internal/svc"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/email/pb/email"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/email.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		email.RegisterEmailServer(grpcServer, server.NewEmailServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
