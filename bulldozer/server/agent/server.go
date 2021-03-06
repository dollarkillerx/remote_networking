package agent

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Client struct {
	Config *Config
	conn   net.Conn
}

// Config is the client configuration
type Config struct {
	Addr     string
	Username string
	Password string
	Verify   func(string, string) bool
}

func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

func (c *Client) Run() {
	defer c.conn.Close()

	reader := bufio.NewReader(c.conn)

	peek, err := reader.Peek(1)
	if err != nil {
		return
	}

	// 判断协议
	switch peek[0] {
	case Version:
	default:
		return
	}

	method := c.chooseMethod()
	if err := WriteMethod(method, c.conn); err != nil || method == MethodNoAcceptable {
		if err != nil {
			log.Printf(`[socks5] write method failed: %s`, err)
		} else {
			log.Printf(`[socks5] methods is not acceptable`)
		}
		return
	}

	if err := method2Handler[method](c, c.conn); err != nil {
		log.Printf(`[socks5] authorization failed: %s`, err)
		return
	}

	// read command
	request, err := ReadRequest(c.conn)
	if err != nil {
		log.Printf(`[socks5] read command failed: %s`, err)
		return
	}

	switch request.Cmd {
	case CmdConnect:
		c.handleConnect(c.conn, request)
	case CmdBind:
		//c.handleBind(c.conn, request)
	case CmdUDP:
		//c.handleUDP(c.conn, request)
	}
}

func (c *Client) chooseMethod() uint8 {
	return MethodNoAuth
	//if supportUserPass {
	//	return MethodUserPass
	//} else if supportNoAuth {
	//	return MethodNoAuth
	//}
	//return MethodNoAcceptable
}

var method2Handler = map[uint8]func(*Client, net.Conn) error{
	MethodNoAuth:   (*Client).authNoAuth,
	MethodUserPass: (*Client).authUserPass,
}

func (c *Client) authNoAuth(conn net.Conn) (err error) {
	return nil
}

func (c *Client) authUserPass(conn net.Conn) (err error) {
	req, err := ReadUserPassRequest(conn)
	if err != nil {
		return
	}

	if !c.Config.Verify(req.Username, req.Password) {
		if e := NewUserPassResponse(UserPassVer, 1).Write(conn); e != nil {
			log.Printf(`[socks5] write reply failed: %s`, e)
		}
		return fmt.Errorf(`verify user %s failed`, req.Username)
	}

	return NewUserPassResponse(UserPassVer, 0).Write(conn)
}
