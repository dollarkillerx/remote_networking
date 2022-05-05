package internal

import (
	"bufio"
	"fmt"
	"log"
	"net"

	pkg2 "github.com/dollarkillerx/remote_networking/pkg"
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

func (g *Gateway) internalServerCore(conn net.Conn) {
	reader := bufio.NewReader(conn)

	pkg := new(pkg2.Package)
	err := pkg.Unpack(reader)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("init success: ", pkg)

	switch pkg.Version[1] {
	case byte(pkg2.PHeartbeat):
		ov := make(chan struct{})
		defer func() {
			close(ov)
		}()

		go func() {
			fmt.Println("new conn")
		loop:
			for {
				select {
				case <-ov:
					break loop
				case data := <-g.mainChannel: // 下发新链接命令
					if data == "new" {
						pg := pkg2.NewPackage(pkg2.PNewConn, []byte(""))
						err := pg.Pack(conn)
						if err != nil {
							log.Println(err)
							continue
						}
						fmt.Println("send new conn")
					}
				}
			}
		}()

		for { // 心跳保活
			scannedPack := new(pkg2.Package)
			err := scannedPack.Unpack(conn)
			if err != nil {
				log.Println(err)
				return
			}

			if pkg.Version[1] == byte(pkg2.PHeartbeat) {
				err := scannedPack.Pack(conn)
				if err != nil {
					log.Println(err)
					return
				}
				continue
			}
		}
	case byte(pkg2.PNewConn): // 新连接放入池中
		g.agentConn <- conn
	}

}

// internalServerCore 处理backend agent 通讯
//func (g *Gateway) internalServerCore(conn net.Conn) {
//	reader := bufio.NewReader(conn)
//
//	scanner := bufio.NewScanner(reader)
//	scanner.Split(pkg2.PackageScannerSplit)
//
//	log.Println("new accept")
//	if scanner.Scan() {
//		log.Println("new accept scan")
//		pkg := new(pkg2.Package)
//		err := pkg.Unpack(bytes.NewReader(scanner.Bytes()))
//		if err != nil {
//			log.Println(err)
//			return
//		}
//
//		fmt.Println("init success: ", pkg)
//
//		switch pkg.Version[1] {
//		case byte(pkg2.PHeartbeat):
//			ov := make(chan struct{})
//			defer func() {
//				close(ov)
//			}()
//
//			go func() {
//				fmt.Println("new conn")
//			loop:
//				for {
//					select {
//					case <-ov:
//						break loop
//					case data := <-g.mainChannel: // 下发新链接命令
//						if data == "new" {
//							pg := pkg2.NewPackage(pkg2.PNewConn, []byte(""))
//							err := pg.Pack(conn)
//							if err != nil {
//								log.Println(err)
//								continue
//							}
//							fmt.Println("send new conn")
//						}
//					}
//				}
//			}()
//			for scanner.Scan() { // 心跳保活
//				scannedPack := new(pkg2.Package)
//				err := scannedPack.Unpack(bytes.NewReader(scanner.Bytes()))
//				if err != nil {
//					log.Println(err)
//					continue
//				}
//
//				if pkg.Version[1] == byte(pkg2.PHeartbeat) {
//					err := scannedPack.Pack(conn)
//					if err != nil {
//						log.Println(err)
//					}
//
//					continue
//				}
//				if err := scanner.Err(); err != nil {
//					log.Println(err)
//					return
//				}
//			}
//
//			if err := scanner.Err(); err != nil {
//				return
//			}
//		case byte(pkg2.PNewConn): // 新连接放入池中
//			g.agentConn <- conn
//		}
//	} else {
//		fmt.Println("aaa")
//		g.agentConn <- conn
//	}
//
//}
