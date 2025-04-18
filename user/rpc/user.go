package main

import (
	"code-hikari/user/rpc/internal/config"
	"code-hikari/user/rpc/internal/server"
	"code-hikari/user/rpc/internal/svc"
	userService "code-hikari/user/rpc/server"
	"flag"
	"fmt"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/nacos"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		userService.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	//api文件生成的默认代码
	defer s.Stop()

	s.AddUnaryInterceptors()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)

	// register service to consul
	err := consul.RegisterService(c.ListenOn, c.Consul)
	if err != nil {
		panic("向consul注册失败:" + err.Error())
	}
	s.Start()

}
