package agent

import (
	"bytes"
	"crypto/tls"
	"io"
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

func (c *Client) handleBind(conn net.Conn, req *Request) {
	log.Printf(`[socks5] "bind" dial server to bind %s for %s`, req.Addr, conn.RemoteAddr())

	listener, err := net.ListenTCP("tcp", nil)
	if err != nil {
		log.Printf(`[socks5] "bind" dial server failed: %s`, err)
		if err := NewReply(HostUnreachable, nil); err != nil {
			log.Printf(`[socks5] "bind" write reply failed: %s`, err)
		}
		return
	}

	// first response: send listen address
	addr, _ := NewAddrFromAddr(listener.Addr(), conn.LocalAddr())
	if err := NewReply(Succeeded, addr).Write(conn); err != nil {
		listener.Close()
		log.Printf(`[socks5] "bind" write reply failed %s`, err)
		return
	}

	newConn, err := listener.AcceptTCP()
	listener.Close()
	if err != nil {
		log.Printf(`[socks5] "bind" bind failed on accept: %s`, err)
		if err := NewReply(Failure, nil).Write(conn); err != nil {
			log.Printf(`[socks5] "bind" write reply failed %s`, err)
		}
		return
	}
	defer newConn.Close()

	// second response: accepted address
	raddr, _ := NewAddr(newConn.RemoteAddr().String())
	if err := NewReply(Succeeded, raddr).Write(conn); err != nil {
		log.Printf(`[socks5] "bind" write reply failed %s`, err)
		return
	}

	log.Printf(`[socks5] "bind" tunnel established %s <-> ?%s`, conn.RemoteAddr(), req.Addr)
	if err := utils.Transport(conn, newConn); err != nil {
		log.Printf(`[socks5] Transport failed: %s`, err)
	}
	log.Printf(`[socks5] "bind" tunnel disconnected %s >-< ?%s`, conn.RemoteAddr(), req.Addr)
}

func (c *Client) handleUDP(conn net.Conn, req *Request) {
	log.Printf(`[socks5] "udp" associate UDP for %s`, conn.RemoteAddr())
	udp, err := net.ListenUDP("udp", nil)
	if err != nil {
		log.Printf(`[socks5] "udp-over-tcp" UDP associate failed on listen: %s`, err)
		if err := NewReply(Failure, nil).Write(conn); err != nil {
			log.Printf(`[socks5] "udp-over-tcp" write reply failed %s`, err)
		}
		return
	}
	defer udp.Close()

	addr, _ := NewAddrFromAddr(udp.LocalAddr(), conn.LocalAddr())
	if err := NewReply(Succeeded, addr).Write(conn); err != nil {
		log.Printf(`[socks5] "udp-over-tcp" write reply failed %s`, err)
		return
	}

	log.Printf(`[socks5] "udp" tunnel established (UDP)%s`, udp.LocalAddr())
	go tunnelUDP(udp, conn)
	if err := waiting4EOF(conn); err != nil {
		log.Printf(`[socks5] "udp" waiting for EOF failed: %s`, err)
	}
	log.Printf(`[socks5] "udp" tunnel disconnected (UDP)%s`, udp.LocalAddr())
}

func tunnelUDP(udp net.PacketConn, conn net.Conn) error {
	errc := make(chan error, 2)
	var clientAddr net.Addr

	go func() {
		b := utils.LPool.Get().([]byte)
		defer utils.LPool.Put(b)

		for {
			n, addr, err := udp.ReadFrom(b)
			if err != nil {
				errc <- err
				return
			}

			dgram, err := ReadUDPDatagram(bytes.NewReader(b[:n]))
			if err != nil {
				errc <- err
				return
			}
			if clientAddr == nil {
				clientAddr = addr
			}
			dgram.Header.Rsv = uint16(len(dgram.Data))
			if err := dgram.Write(conn); err != nil {
				errc <- err
				return
			}
		}
	}()

	go func() {
		for {
			dgram, err := ReadUDPDatagram(conn)
			if err != nil {
				errc <- err
				return
			}

			if clientAddr == nil {
				continue
			}
			dgram.Header.Rsv = 0
			buf := bytes.NewBuffer(nil)
			dgram.Write(buf)
			if _, err := udp.WriteTo(buf.Bytes(), clientAddr); err != nil {
				errc <- err
				return
			}
		}
	}()

	return <-errc
}

func waiting4EOF(conn net.Conn) (err error) {
	b := utils.SPool.Get().([]byte)
	defer utils.SPool.Put(b)
	for {
		_, err = conn.Read(b)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
	}
	return
}
