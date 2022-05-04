package internal

import (
	"bufio"
	"bytes"
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

	scanner := bufio.NewScanner(reader)
	scanner.Split(PackageScannerSplit)

	pkg := new(Package)
	if scanner.Scan() {
		err := pkg.Unpack(bytes.NewReader(scanner.Bytes()))
		if err != nil {
			log.Println(err)
			return
		}

		switch pkg.Version[1] {
		case byte(PHeartbeat):
			ov := make(chan struct{})
			defer func() {
				close(ov)
			}()

			go func() {
			loop:
				for {
					select {
					case <-ov:
						break loop
					case data := <-g.mainChannel: // 下发新链接命令
						if data == "new" {
							pg := NewPackage(PNewConn, []byte(""))
							err := pg.Pack(conn)
							if err != nil {
								log.Println(err)
								continue
							}
						}
					}
				}
			}()
			for scanner.Scan() { // 心跳保活
				scannedPack := new(Package)
				err := scannedPack.Unpack(bytes.NewReader(scanner.Bytes()))
				if err != nil {
					log.Println(err)
					continue
				}
				if pkg.Version[1] == byte(PHeartbeat) {
					continue
				}
				if err := scanner.Err(); err != nil {
					log.Println(err)
					return
				}
			}

			if err := scanner.Err(); err != nil {
				return
			}
		case byte(PNewConn): // 新连接放入池中
			g.agentConn <- conn
		}
	}
}
