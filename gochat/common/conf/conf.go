<<<<<<< HEAD:gochat/conf/conf.go
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
	var file string = "../conf/conf.json"
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(f).Decode(&ConfigureInfo)
	if err != nil {
		panic(err)
	}
=======
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
>>>>>>> 65888be0467c07f28b57cd4073746ad38fcc0782:gochat/common/conf/conf.go
}