package utils

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dollarkillerx/remote_networking/utils"
)

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
