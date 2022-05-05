package main

import (
	"github.com/dollarkillerx/remote_networking/bulldozer/conf"
	"github.com/dollarkillerx/remote_networking/bulldozer/server/agent"

	"log"
	"net"
)

func main() {
	err := conf.InitAgentConfig()
	if err != nil {
		log.Fatalln(err)
	}

	log.SetFlags(log.Llongfile)

	// 本地建立socks5
	listener, err := net.Listen("tcp", conf.AgentConfig.LocalAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("listen run: ", conf.AgentConfig.LocalAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		client := agent.NewClient(conn)
		go client.Run()
	}
}
