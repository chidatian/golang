package conf

import (
	"encoding/json"
	"os"
)

type (
	Configure struct {
		Server serverInfo `json:"ServerInfo"`
		Redis redisInfo `json:"RedisInfo"`
	}

	serverInfo struct {
		Host string `json:"Host"`
	}

	redisInfo struct {
		Host string `json:"Host"`
	}
)

var ConfigureInfo Configure

func init() {
	var file string = "./conf/conf.json"
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(f).Decode(&ConfigureInfo)
	if err != nil {
		panic(err)
	}
}