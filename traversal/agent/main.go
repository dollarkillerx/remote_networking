package main

import (
	"log"
	"net"
	"time"

	"github.com/dollarkillerx/remote_networking/pkg"
	"github.com/dollarkillerx/remote_networking/traversal/conf"
)

var internalGatewayAddr *net.TCPAddr

func main() {
	// 1. 一个主链接控制 负责子链接的控制

	conf.InitConf()

	addr, err := net.ResolveTCPAddr("tcp4", conf.Conf.InternalGateway)
	if err != nil {
		panic(err)
	}
	internalGatewayAddr = addr

	log.Printf("Agent InternalGateway: %s LocalProxy: %s \n", conf.Conf.InternalGateway, conf.Conf.AgentProxyAddr)
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
			pg := pkg.NewPackage(pkg.PHeartbeat, []byte(conf.Conf.Token))
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
	local, err := net.Dial("tcp", conf.Conf.AgentProxyAddr)
	if err != nil {
		log.Println(err)
		return
	}

	remove, err := net.DialTCP("tcp", nil, internalGatewayAddr)
	if err != nil {
		log.Println(err)
		return
	}

	pk := pkg.NewPackage(pkg.PNewConn, []byte(conf.Conf.Token))
	err = pk.Pack(remove)
	if err != nil {
		log.Println(err)
		return
	}

	//log.Println("new conn")

	pkg.Transport(remove, local)
}
