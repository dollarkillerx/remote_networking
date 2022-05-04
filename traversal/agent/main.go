package main

import (
	"bufio"
	"bytes"
	"log"
	"net"
	"time"

	"github.com/dollarkillerx/remote_networking/pkg"
)

var internalGatewayAddr *net.TCPAddr

func main() {
	// 1. 一个主链接控制 负责子链接的控制
	addr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8082")
	if err != nil {
		panic(err)
	}
	internalGatewayAddr = addr

	mainCore()

}

func mainCore() error {
	var tcpConn *net.TCPConn
	var err error
	tcpConn, err = net.DialTCP("tcp", nil, internalGatewayAddr)
	if err != nil {
		return err
	}

	go func() {
		for {
			time.Sleep(time.Second)
			pg := pkg.NewPackage(pkg.PHeartbeat, []byte(""))
			err := pg.Pack(tcpConn)
			if err != nil {
				log.Println(err)
				break
			}
		}
	}()

	scanner := bufio.NewScanner(tcpConn)
	scanner.Split(pkg.PackageScannerSplit)
	for scanner.Scan() {
		scannedPack := new(pkg.Package)
		err := scannedPack.Unpack(bytes.NewReader(scanner.Bytes()))
		if err != nil {
			log.Println(err)
			break
		}

		if scannedPack.Version[1] == byte(pkg.PNewConn) {

		}
	}

	return scanner.Err()
}
