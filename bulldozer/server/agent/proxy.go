package agent

import (
	"crypto/tls"
	"log"
	"net"
	"strings"

	"github.com/dollarkillerx/remote_networking/bulldozer/conf"
	"github.com/dollarkillerx/remote_networking/bulldozer/pkg"
	"github.com/dollarkillerx/remote_networking/bulldozer/utils"
)

func (c *Client) handleConnect(conn net.Conn, req *Request) {
	var nextHop net.Conn
	var err error
	var isProxy bool // pac 更具DNS or 域名匹配选择是否代理
	isProxy = true

	if isProxy {
		nextHop, err = c.dialServer(req.Addr.String())
		if err != nil {
			log.Printf(`[socks5] "connect" dial server failed: %s`, err)
			if err = NewReply(HostUnreachable, nil).Write(conn); err != nil {
				log.Printf(`[socks5] "connect" write reply failed: %s`, err)
			}
			return
		}
		defer nextHop.Close()
	} else {
		nextHop, err = net.Dial("tcp", req.Addr.String())
		if err != nil {
			if err = NewReply(HostUnreachable, nil).Write(conn); err != nil {
				log.Printf(`[socks5] "connect" write reply failed: %s`, err)
			}
			return
		}
		defer nextHop.Close()
	}

	// 确认代理成功
	if err = NewReply(Succeeded, nil).Write(conn); err != nil {
		log.Printf(`[socks5] "connect" write reply failed: %s`, err)
		return
	}

	if err := utils.Transport(conn, nextHop); err != nil {
		log.Printf(`[socks5] "connect" transport failed: %s`, err)
	}
}

func (c *Client) dialServer(addr string) (net.Conn, error) {
	dial, err := tls.Dial("tcp", conf.AgentConfig.BackendAddr, nil)
	if err != nil {
		log.Fatalln("远程服务器链接失败: ", err)
	}

	divert := utils.NewDivert(conf.AgentConfig.Token)
	key := divert.NewKey()

	header := strings.ReplaceAll(rb, "TMG", key)

	dial.Write([]byte(header))

	var bull pkg.Bulldozer
	bull.ConnType = pkg.TcpType
	bull.Addr = addr

	dial.Write(append(bull.ToJSON(), '\n'))

	return dial, nil
}

var rb = "GET / HTTP/1.1\nHost: m.baidu.com\nUser-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:99.0) Gecko/20100101 Firefox/99.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8\nAccept-Language: zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2\nAccept-Encoding: gzip, deflate\nConnection: close\nCookie: BIDUPSID=538DF3B8B722AA6B57889828A36D9675; PSTM=1602731420; BAIDUID=9A382B3A2ED31BAC9FEEED65B583E825:FG=1; H_PS_PSSID=36309_31660_34813_35912_36167_34584_35978_36278_36232_26350_36302_36313_36061; BA_HECTOR=ag2l04a58l8h0g257l1h76ifl0r\nUpgrade-Insecure-Requests: 1\nPragma: no-cache\nAuthorization: TMG\nCache-Control: no-cache\r\n\r\n"
