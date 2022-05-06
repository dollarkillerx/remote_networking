package main

import (
	"github.com/dollarkillerx/remote_networking/traversal/backend/internal"
	"github.com/dollarkillerx/remote_networking/traversal/conf"

	"log"
)

func main() {
	// 1. 我们需要两个入口
	// 1.1 服务网关 (外部服务调用访问我们内网服务的网关)
	// 1.2 内部服务入口 (内网服务器与网关通信的入口)

	conf.InitConf()

	gateway, err := internal.NewGateway(conf.Conf.ExternalGateway, conf.Conf.InternalGateway)
	if err != nil {
		panic(err)
	}

	log.Printf("ExternalGateawy: %s InternalGateawy: %s \n", conf.Conf.ExternalGateway, conf.Conf.InternalGateway)
	if err := gateway.Run(); err != nil {
		panic(err)
	}
}
