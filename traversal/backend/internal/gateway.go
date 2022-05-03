package internal

import "net"

type Gateway struct {
	externalGateway *net.TCPListener // 服务网关 (外部服务调用访问我们内网服务的网关)
	internalGateway *net.TCPListener // 内部服务入口 (内网服务器与网关通信的入口)
}

func NewGateway(externalGateway string, internalGateway string) (*Gateway, error) {
	externalGatewayIP, err := net.ResolveTCPAddr("tcp", externalGateway)
	if err != nil {
		return nil, err
	}

	internalGatewayIP, err := net.ResolveTCPAddr("tcp", internalGateway)
	if err != nil {
		return nil, err
	}

	externalGatewayTcp, err := net.ListenTCP("tcp", externalGatewayIP)
	if err != nil {
		return nil, err
	}

	internalGatewayTcp, err := net.ListenTCP("tcp", internalGatewayIP)
	if err != nil {
		return nil, err
	}

	return &Gateway{
		externalGateway: externalGatewayTcp,
		internalGateway: internalGatewayTcp,
	}, nil
}

func (g *Gateway) Run() error {
	go g.internalServer()
	return nil
}
