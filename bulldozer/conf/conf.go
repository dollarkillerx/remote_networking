package conf

import (
	"encoding/json"
	"io/ioutil"
)

type agentConfig struct {
	LocalAddr   string `json:"local_addr"`
	Token       string `json:"token"`
	BackendAddr string `json:"backend_addr"`
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
