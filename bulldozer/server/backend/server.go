package backend

import (
	"bufio"
	"fmt"
	"github.com/dollarkillerx/remote_networking/bulldozer/conf"
	"github.com/dollarkillerx/remote_networking/bulldozer/utils"
	"log"
	"net"
	"strings"
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
			fmt.Println("TCP divert................")
		}
	}

	html, err := getTmpHTML()
	if err != nil {
		log.Println(err)
		return
	}

	conn.Write(html)
}
