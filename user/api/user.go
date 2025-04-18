package main

import (
	"code-hikari/common-go"
	"code-hikari/user/api/internal/config"
	"code-hikari/user/api/internal/handler"
	"code-hikari/user/api/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"strconv"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(common.ErrorHandler)
	httpx.SetOkHandler(common.SuccessHandler)

	// register service to consul
	err := consul.RegisterService(c.Host+":"+strconv.Itoa(c.Port), c.Consul)
	if err != nil {
		panic("向consul注册失败:" + err.Error())
	}

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
