package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/dollarkillerx/remote_networking/utils"
)

var (
	SPool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 576)
		},
	} // small buff pool
	LPool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 64*1024+262)
		},
	} // large buff pool for udp
)

// Transport rw1 and rw2
func Transport(rw1, rw2 io.ReadWriter) error {
	errc := make(chan error, 1)
	go func() {
		b := LPool.Get().([]byte)
		defer LPool.Put(b)

		_, err := io.CopyBuffer(rw1, rw2, b)
		errc <- err
	}()

	go func() {
		b := LPool.Get().([]byte)
		defer LPool.Put(b)

		_, err := io.CopyBuffer(rw2, rw1, b)
		errc <- err
	}()

	if err := <-errc; err != nil && err != io.EOF {
		return err
	}
	return nil
}

type Divert struct {
	Timeout int64  `json:"timeout"`
	Key     []byte `json:"-"`
}

func NewDivert(key string) Divert {
	return Divert{Key: []byte(key)}
}

func (d Divert) NewKey() string {
	d.Timeout = time.Now().Add(time.Hour * 5).Unix()
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
