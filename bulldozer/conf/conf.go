package conf

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

type agentConfig struct {
	LocalAddr   string   `json:"local_addr"`
	Token       string   `json:"token"`
	BackendAddr string   `json:"backend_addr"`
	Pac         bool     `json:"pac"`
	Dns         string   `json:"dns"`
	ProxyList   []string `json:"proxy_list"`
	NoProxyList []string `json:"no_proxy_list"`
}

func (a *agentConfig) GetDNS() string {
	if a.Dns == "" {
		return "8.8.8.8:53"
	}
	return a.Dns
}

type backendConfig struct {
	ListenAddr string `json:"listen_addr"`
	Token      string `json:"token"`
	CertFile   string `json:"cert_file"`
	KeyFile    string `json:"key_file"`
}

var AgentConfig *agentConfig

var BackendConfig *backendConfig

func InitAgentConfig() error {
	file, err := ioutil.ReadFile("agent.json")
	if err != nil {
		var agentConf agentConfig
		agentConf.NoProxyList = []string{}
		agentConf.ProxyList = []string{}
		marshal, err := json.Marshal(agentConf)
		if err != nil {
			return err
		}

		var str bytes.Buffer
		_ = json.Indent(&str, marshal, "", "    ")
		ioutil.WriteFile("agent.json", str.Bytes(), 00666)
		return err
	}

	var ag agentConfig
	err = json.Unmarshal(file, &ag)
	if err != nil {
		panic(err)
	}

	AgentConfig = &ag
	return nil
}

func InitBackendConfig() error {
	file, err := ioutil.ReadFile("backend.json")
	if err != nil {
		var backendC backendConfig
		marshal, err := json.Marshal(backendC)
		if err != nil {
			return err
		}

		var str bytes.Buffer
		_ = json.Indent(&str, marshal, "", "    ")
		ioutil.WriteFile("backend.json", str.Bytes(), 00666)
		return err
	}

	var ag backendConfig
	err = json.Unmarshal(file, &ag)
	if err != nil {
		panic(err)
	}

	BackendConfig = &ag
	return nil
}
