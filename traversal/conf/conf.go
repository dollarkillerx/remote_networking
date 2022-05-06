package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type traversalConf struct {
	Token           string `json:"token"`
	AgentProxyAddr  string `json:"agent_proxy_addr"` // agent代理服务地址
	ExternalGateway string `json:"external_gateway"` // 外部访问地址
	InternalGateway string `json:"internal_gateway"` // 内部代理地址
}

var Conf *traversalConf

func InitConf() {
	file, err := ioutil.ReadFile("traversal.json")
	if err != nil {
		log.Fatalln("read config error: ", err)
	}

	var conf traversalConf
	err = json.Unmarshal(file, &conf)
	if err != nil {
		log.Fatalln("pares config error: ", err)
	}

	Conf = &conf
}
