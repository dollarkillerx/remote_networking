package utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDns(t *testing.T) {
	dns, err := SearchDns("www.baidu.com", "8.8.8.8:53")
	if err != nil {
		return
	}

	marshal, err := json.Marshal(dns)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(marshal))
}
