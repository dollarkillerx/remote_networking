package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"strings"

	"github.com/dollarkillerx/remote_networking/bulldozer/conf"
	"github.com/dollarkillerx/remote_networking/bulldozer/utils"
)

func main() {
	err := conf.InitAgentConfig()
	if err != nil {
		log.Fatalln(err)
	}

	dial, err := tls.Dial("tcp", conf.AgentConfig.BackendAddr, nil)
	if err != nil {
		panic(err)
	}
	defer dial.Close()

	divert := utils.NewDivert(conf.AgentConfig.Token)
	key := divert.NewKey()

	fmt.Println(key)
	header := strings.ReplaceAll(rb, "TMG", key)

	dial.Write([]byte(header))
}

var rb = "GET / HTTP/1.1\nHost: m.baidu.com\nUser-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:99.0) Gecko/20100101 Firefox/99.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8\nAccept-Language: zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2\nAccept-Encoding: gzip, deflate\nConnection: close\nCookie: BIDUPSID=538DF3B8B722AA6B57889828A36D9675; PSTM=1602731420; BAIDUID=9A382B3A2ED31BAC9FEEED65B583E825:FG=1; H_PS_PSSID=36309_31660_34813_35912_36167_34584_35978_36278_36232_26350_36302_36313_36061; BA_HECTOR=ag2l04a58l8h0g257l1h76ifl0r\nUpgrade-Insecure-Requests: 1\nPragma: no-cache\nAuthorization: TMG\nCache-Control: no-cache\r\n\r\n"
