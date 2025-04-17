package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zhihu/application/applet/internal/config"
	"go-zhihu/application/applet/internal/handler"
	"go-zhihu/application/applet/internal/svc"
	"go-zhihu/common"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/applet-api.yaml", "the config file")

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

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
