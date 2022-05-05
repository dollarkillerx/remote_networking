package main

import (
	"crypto/tls"
	"log"

	"github.com/dollarkillerx/remote_networking/bulldozer/conf"
	"github.com/dollarkillerx/remote_networking/bulldozer/server/backend"
)

func main() {
	err := conf.InitBackendConfig()
	if err != nil {
		log.Fatalln(err)
	}

	cer, err := tls.LoadX509KeyPair(conf.BackendConfig.CertFile, conf.BackendConfig.KeyFile)
	if err != nil {
		log.Fatalln(err)
	}

	tlsCfg := &tls.Config{Certificates: []tls.Certificate{cer}}

	listen, err := tls.Listen("tcp", conf.BackendConfig.ListenAddr, tlsCfg)
	if err != nil {
		panic(err)
	}

	log.Println("listen run: ", conf.BackendConfig.ListenAddr)

	for {
		accept, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go backend.Server(accept)
	}
}
