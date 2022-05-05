package tj

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"testing"
)

func TestHttp(t *testing.T) {
	dial, err := net.Dial("tcp", "m.baidu.com:80")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = dial.Write([]byte(htp))
	if err != nil {
		panic(err)
	}

	var body []byte

	reader := bufio.NewReader(dial)
	for {
		var buf = make([]byte, 512)
		i, err := reader.Read(buf)
		if err != nil {
			break
		}

		body = append(body, buf[:i]...)
	}

	fmt.Println(string(body))
	ioutil.WriteFile("aaa.body", body, 00666)
}

func TestHttpServer(t *testing.T) {
	log.SetFlags(log.Llongfile)
	conn, err := net.Listen("tcp", "0.0.0.0:8175")
	if err != nil {
		panic(err)
	}

	for {
		conn, _ := conn.Accept()
		go func() {
			defer conn.Close()
			reader := bufio.NewReader(conn)

			var header string
			for {
				readByte, err := reader.ReadBytes('\n')
				if err != nil {
					log.Println(err)
					break
				}
				if strings.TrimSpace(string(readByte)) == "" {
					break
				}
				header += string(readByte)
			}
			fmt.Println(header)

			file, err := ioutil.ReadFile("aaa.body")
			if err != nil {
				panic(err)
			}

			conn.Write(file)
		}()
	}

	for {
		conn, _ := conn.Accept()
		go func() {
			defer conn.Close()
			reader := bufio.NewReader(conn)

			var header string
			for {
				readByte, err := reader.ReadBytes('\n')
				if err != nil {
					log.Println(err)
					break
				}
				if strings.TrimSpace(string(readByte)) == "" {
					break
				}
				header += string(readByte)
			}
			fmt.Println(header)

			conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
			conn.Write([]byte("Content-Length: 2\r\n"))
			conn.Write([]byte("Content-Type:text/html:charset=UTF-8\r\n\r\n"))
			conn.Write([]byte("ok"))
			fmt.Println("over")
		}()
	}
}

var htp = "GET / HTTP/1.1\nHost: m.baidu.com\nUser-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:99.0) Gecko/20100101 Firefox/99.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8\nAccept-Language: zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2\nAccept-Encoding: gzip, deflate\nConnection: close\nCookie: BIDUPSID=538DF3B8B722AA6B57889828A36D9675; PSTM=1602731420; BAIDUID=9A382B3A2ED31BAC9FEEED65B583E825:FG=1; H_PS_PSSID=36309_31660_34813_35912_36167_34584_35978_36278_36232_26350_36302_36313_36061; BA_HECTOR=ag2l04a58l8h0g257l1h76ifl0r\nUpgrade-Insecure-Requests: 1\nPragma: no-cache\nCache-Control: no-cache\r\n\r\n"