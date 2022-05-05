package main

import (
	"io"
	"log"
	"net"
	"time"

	"github.com/dollarkillerx/remote_networking/pkg"
)

var internalGatewayAddr *net.TCPAddr

func main() {
	// 1. 一个主链接控制 负责子链接的控制
	addr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8182")
	if err != nil {
		panic(err)
	}
	internalGatewayAddr = addr

	// 断线重连
	for {
		err := mainCore()
		if err != nil {
			log.Println(err)
		}

		time.Sleep(time.Second)
	}

}

func mainCore() error {
	var tcpConn *net.TCPConn
	var err error
	tcpConn, err = net.DialTCP("tcp", nil, internalGatewayAddr)
	if err != nil {
		return err
	}

	// 主链接心跳保活
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

	log.Println("init agent success")

	for {
		scannedPack := new(pkg.Package)
		err := scannedPack.Unpack(tcpConn)
		if err != nil {
			log.Println(err)
			return err
		}

		switch scannedPack.Version[1] {
		case byte(pkg.PHeartbeat):
		case byte(pkg.PNewConn):
			go newConn()
		}
	}

	return nil
}

func newConn() {
	local, err := net.Dial("tcp", "127.0.0.1:5432")
	if err != nil {
		log.Println(err)
		return
	}

	remove, err := net.DialTCP("tcp", nil, internalGatewayAddr)
	if err != nil {
		log.Println(err)
		return
	}

	pkg := pkg.NewPackage(pkg.PNewConn, []byte(""))
	err = pkg.Pack(remove)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("new conn")

	go ioCopy(remove, local)
	ioCopy(local, remove)
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
			log.Println(err)
			break
		}
	}
}
