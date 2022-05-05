package tj

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/dollarkillerx/remote_networking/utils"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"testing"
	"time"
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

func TestHttpTLSServer(t *testing.T) {
	// openssl genrsa -out server.key 2048
	// openssl ecparam -genkey -name secp384r1 -out server.key
	// openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

	log.SetFlags(log.Llongfile)

	cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Println(err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}

	conn, err := tls.Listen("tcp", "0.0.0.0:8175", config)
	if err != nil {
		panic(err)
	}

	// https://0.0.0.0:8175/

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

type Divert struct {
	Timeout int64  `json:"timeout"`
	Key     []byte `json:"-"`
}

func NewDivert(key string) Divert {
	return Divert{Key: []byte(key)}
}

func (d Divert) NewKey() string {
	d.Timeout = time.Now().Add(time.Second * 5).Unix()
	marshal, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	encrypt, err := utils.AesEncrypt(marshal, d.Key)
	if err != nil {
		panic(err)
	}

	return utils.Base64Encode(encrypt)
}

func (d Divert) Check(key string) bool {
	decode, err := utils.Base64Decode(key)
	if err != nil {
		return false
	}

	decrypt, err := utils.AesDecrypt(decode, d.Key)
	if err != nil {
		return false
	}

	err = json.Unmarshal(decrypt, &d)
	if err != nil {
		return false
	}

	fmt.Println(d.Timeout)
	if d.Timeout < time.Now().Unix() {
		return false
	}

	return true
}

func TestDivert(t *testing.T) {
	key := "378def8265924fbb80d9b5f79a912e2d"

	divert := NewDivert(key)
	pKey := divert.NewKey()
	fmt.Println(pKey)

	time.Sleep(time.Second * 6)
	fmt.Println(divert.Check(pKey))
}
