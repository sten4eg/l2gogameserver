package config

import (
	"encoding/json"
	"l2gogameserver/data/logger"
	"os"
)

type Config struct {
	GameServer   GameServer `json:"gameserver"`
	isConfigInit bool
}
type GameServer struct {
	ServerId           int          `json:"serverId"`
	AcceptAlternateId  bool         `json:"acceptAlternateId"`
	ReserveHostOnLogin bool         `json:"reserveHostOnLogin"`
	Port               int16        `json:"port"`
	ServerListBrackets bool         `json:"serverListBrackets"`
	GMOnly             bool         `json:"GMOnly"`
	ServerListAge      byte         `json:"serverListAge"`
	ServerListType     string       `json:"serverListType"`
	MaxPlayer          int          `json:"maxPlayer"`
	HexId              []byte       `json:"hexId"`
	PortForLS          string       `json:"portForLS"`
	Database           DatabaseType `json:"database"`
	Debug              Debug        `json:"debug"`
}
type DatabaseType struct {
	Name         string `json:"name"`
	Host         string `json:"host"`
	Schema       string `json:"schema"`
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

var configInstance Config

const MaxAdena = 99_900_000_000

func Get() GameServer {
	if !configInstance.isConfigInit {
		read()
	}
	return configInstance.GameServer
}

// todo сделать так чтобы read не возвращал Config
func read() Config {
	file, err := os.Open("./config/config.json")
	if err != nil {
		logger.Error.Panicln("Failed to load /config/config.json file")
	}

	decoder := json.NewDecoder(file)
	var conf Config
	err = decoder.Decode(&conf)
	if err != nil {
		logger.Error.Panicln("Failed to decode config file")
	}
	conf.isConfigInit = true
	configInstance = conf

	return configInstance
}
func LoadAllConfig() {
	read()
}

func GetHexId() []byte {
	return configInstance.GameServer.HexId
}
func GetLoginServerPort() string {
	return configInstance.GameServer.PortForLS
}

func GetServerId() int {
	return configInstance.GameServer.ServerId
}

func GetAcceptAlternateId() bool {
	return configInstance.GameServer.AcceptAlternateId
}

func GetReserveHostOnLogin() bool {
	return configInstance.GameServer.ReserveHostOnLogin
}

func GetPort() int16 {
	return configInstance.GameServer.Port
}

func GetMaxPlayer() int {
	return configInstance.GameServer.MaxPlayer
}

func GetServerListBrackets() bool {
	return configInstance.GameServer.ServerListBrackets
}
func GetGMOnly() bool {
	return configInstance.GameServer.GMOnly
}
func GetServerListAge() byte {
	return configInstance.GameServer.ServerListAge
}
func GetServerListType() string {
	return configInstance.GameServer.ServerListType
}
