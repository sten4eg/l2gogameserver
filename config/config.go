package config

import (
	"encoding/json"
	"os"
)

type Data struct {
	GameServer GameServer `json:"gameserver"`
}
type GameServer struct {
	Database DatabaseType `json:"database"`
	Debug    Debug        `json:"debug"`
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
type Debug struct {
	ShowPackets      bool `json:"show_packets"`
	EnableNPC        bool `json:"enable_load_npc"`
	EnabledSkills    bool `json:"enabled_load_skills"`
	EnabledItems     bool `json:"enabled_items"`
	EnabledSpawnlist bool `json:"enabled_spawnlist"`
}

var config Data

func Get() GameServer {
	if (config == Data{}) {
		read()
	}
	return config.GameServer
}

func read() Data {
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

/*
	Загрузка конфигурации
*/
func LoadAllConfig() {
	read()
}

//func LoadDebug() Data {
//	var config Data
//	file, err := os.Open("./config/debug.json")
//	if err != nil {
//		panic("Failed to load config file")
//	}
//
//	decoder := json.NewDecoder(file)
//	err = decoder.Decode(&config)
//	if err != nil {
//		panic("Failed to decode config file")
//	}
//	return config
//}
