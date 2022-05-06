package ip2region

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestIp2region(t *testing.T) {
	region, err := New("../../ip2region.db")
	defer region.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	ip, err := region.MemorySearch("192.168.88.11")
	fmt.Println(ip, err)
	marshal, err := json.Marshal(ip)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))

	//ip.City = "内网IP"

}
