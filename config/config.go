package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	LoginServer LoginServerType
	GameServers []GameServerType
}

type GameServerConfigObject struct {
	LoginServer LoginServerType
	GameServer  GameServerType
}

type DatabaseType struct {
	Name     string
	Host     string
	Port     uint16
	User     string
	Password string
}

type CacheType struct {
	Host     string
	Port     int
	Password string
}

type LoginServerType struct {
	Host       string
	AutoCreate bool
	Database   DatabaseType
}

type GameServerType struct {
	Name       string
	InternalIP string
	ExternalIP string
	Port       int
	Database   DatabaseType
	Cache      CacheType
	Options    OptionsType
}

type OptionsType struct {
	MaxPlayers uint16
	Testing    bool
}

func Read() Config {

	var config Config
	file, err := os.Open("./config/config.json")
	if err != nil {
		log.Fatal("Failed to load config file")
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Failed to decode config file")
	}
	return config
}
