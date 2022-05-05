package internal

import (
	"io"
	"log"
	"time"
)

func (g *Gateway) externalServer() error {
	for {
		accept, err := g.externalGateway.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// 下发新链接指令
		g.mainChannel <- "new"
		select {
		case <-time.After(time.Second): // 超时控制
			accept.Close()
			break
		case conn := <-g.agentConn:
			go ioCopy(accept, conn)
			ioCopy(conn, accept)
		}
	}
}

func ioCopy(server io.Writer, client io.Reader) {
	for {
		var b [1024]byte
		read, err := client.Read(b[:])
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
			break
		}

		if _, err := server.Write(b[:read]); err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
			break
		}
	}
}
