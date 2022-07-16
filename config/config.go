package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var GConfig *ServerConfig

type ServerConfig struct {
	ServerName string
	IP         string
	Port       int
	Version    string
}

func init() {
	reload()
}

func reload() {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("read config file err")
		return
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("read config file err")
		return
	}

	tmp := &ServerConfig{}
	err = json.Unmarshal(bytes, tmp)
	if err != nil {
		fmt.Println("read config file err")
		return
	}
	GConfig = tmp
}
