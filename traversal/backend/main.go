package main

import "github.com/dollarkillerx/remote_networking/traversal/backend/internal"

func main() {
	// 1. 我们需要两个入口
	// 1.1 服务网关 (外部服务调用访问我们内网服务的网关)
	// 1.2 内部服务入口 (内网服务器与网关通信的入口)

	gateway, err := internal.NewGateway("0.0.0.0:8081", "0.0.0.0:8082")
	if err != nil {
		panic(err)
	}

	if err := gateway.Run(); err != nil {
		panic(err)
	}
}
