package internal

import (
	"bufio"
	"log"
	"net"
)

func (g *Gateway) internalServer() {
	for {
		// 内网服务 链接到 gateway
		accept, err := g.internalGateway.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go g.internalServerCore(accept)
	}
}

// internalServerCore 处理backend agent 通讯
func (g *Gateway) internalServerCore(conn net.Conn) {

	reader := bufio.NewReader(conn)
	for {
		reader.ReadBytes()
	}
}
