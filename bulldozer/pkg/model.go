package pkg

import "encoding/json"

type Bulldozer struct {
	ConnType ConnType `json:"conn_type"`
	Addr     string   `json:"addr"`
}

func (b Bulldozer) ToJSON() []byte {
	marshal, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	return marshal
}

type ConnType uint8

const (
	ConnUndefined ConnType = 0
	TcpType       ConnType = 1
)
