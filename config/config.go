package config

import (
	"encoding/json"
	"l2gogameserver/data/logger"
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
	// EnabledCacheHTML - если false, тогда не будет записываться в кэш, удобно для
	// редактирования HTML диалогов и просмотра в игре при каждом обращении.
	EnabledCacheHTML bool `json:"enabled_cache_html"`
}

var config Data

func Get() GameServer {
	if (config == Data{}) {
		read()
	}
	return config.GameServer
}

func read() Data {
	file, err := os.Open("./config/con1fig.json")
	if err != nil {
		logger.Error.Panicln("Failed to load config file")
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		logger.Error.Panicln("Failed to decode config file")
	}
	return config
}

/*
	Загрузка конфигурации
*/
func LoadAllConfig() {
	read()
}
