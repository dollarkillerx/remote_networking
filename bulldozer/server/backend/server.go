package backend

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/dollarkillerx/remote_networking/bulldozer/conf"
	"github.com/dollarkillerx/remote_networking/bulldozer/pkg"
	"github.com/dollarkillerx/remote_networking/bulldozer/utils"
)

func Server(conn net.Conn) {
	defer func() {
		conn.Close()
	}()

	reader := bufio.NewReader(conn)

	var header string
	var token string
	for {
		readByte, err := reader.ReadBytes('\n')
		if err != nil {
			log.Println(err)
			break
		}

		ph := strings.TrimSpace(string(readByte))

		if strings.Contains(ph, "Authorization") {
			repeat := strings.Split(ph, ":")
			if len(repeat) == 2 {
				token = strings.TrimSpace(repeat[1])
			}
		}

		if ph == "" {
			break
		}
		header += string(readByte)
	}

	if token != "" {
		divert := utils.NewDivert(conf.BackendConfig.Token)
		check := divert.Check(token)
		if check {
			tcpDivert(conn, reader)
		}
	}

	html, err := getTmpHTML()
	if err != nil {
		log.Println(err)
		return
	}

	conn.Write(html)
}

func tcpDivert(conn net.Conn, reader *bufio.Reader) {
	bt, err := reader.ReadBytes('\n')
	if err != nil {
		log.Println(err)
		return
	}

	var bull pkg.Bulldozer
	err = json.Unmarshal(bytes.TrimSpace(bt), &bull)
	if err != nil {
		log.Println(err)
		return
	}

	var nextHop net.Conn

	switch bull.ConnType {
	case pkg.ConnUndefined:
		return
	case pkg.TcpType:
		fmt.Println("tcp: ", bull.Addr)
		nextHop, err = net.Dial("tcp", bull.Addr)
		if err != nil {
			return
		}

		defer nextHop.Close()
	}

	if err := utils.Transport(conn, nextHop); err != nil {
		log.Println(err)
	}
}
