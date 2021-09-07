package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	GameServer GameServer `json:"gameserver"`
}
type GameServer struct {
	Database DatabaseType `json:"database"`
}
type DatabaseType struct {
	Name         string `json:"name"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	SSLMode      string `json:"sslmode"`
	PoolMaxConns string `json:"pool_max_conns"`
}

func Read() Config {

	var config Config
	file, err := os.Open("./config/config.json")
	if err != nil {
		panic("Failed to load config file")
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic("Failed to decode config file")
	}
	return config
}
