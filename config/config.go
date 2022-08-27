package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type config struct {
	GameServer   GameServer `yaml:"gameserver"`
	isConfigInit bool
}
type GameServer struct {
	ServerId           int            `yaml:"serverId"`
	AcceptAlternateId  bool           `yaml:"acceptAlternateId"`
	ReserveHostOnLogin bool           `yaml:"reserveHostOnLogin"`
	Port               int16          `yaml:"port"`
	ServerListBrackets bool           `yaml:"serverListBrackets"`
	GMOnly             bool           `yaml:"GMOnly"`
	ServerListAge      byte           `yaml:"serverListAge"`
	ServerListType     string         `yaml:"serverListType"`
	MaxPlayer          int            `yaml:"maxPlayer"`
	HexId              []byte         `yaml:"hexId"`
	PortForLS          string         `yaml:"portForLS"`
	Database           DatabaseConfig `yaml:"database"`
	Debug              Debug          `yaml:"debug"`
}
type DatabaseConfig struct {
	Name         string `yaml:"name"`
	Host         string `yaml:"host"`
	Schema       string `yaml:"schema"`
	Port         string `yaml:"port"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	SSLMode      string `yaml:"sslmode"`
	PoolMaxConns string `yaml:"pool_max_conns"`
}
type Debug struct {
	ShowPackets      bool `yaml:"show_packets"`
	EnableNPC        bool `yaml:"enable_load_npc"`
	EnabledSkills    bool `yaml:"enabled_load_skills"`
	EnabledItems     bool `yaml:"enabled_items"`
	EnabledSpawnlist bool `yaml:"enabled_spawnlist"`
	// EnabledCacheHTML - если false, тогда не будет записываться в кэш, удобно для
	// редактирования HTML диалогов и просмотра в игре при каждом обращении.
	EnabledCacheHTML bool `yaml:"enabled_cache_html"`
}

var configInstance config

const MaxAdena = 99_900_000_000

//func Get() GameServer {
//	if !configInstance.isConfigInit {
//		read()
//	}
//	return configInstance.GameServer
//}

func read() error {
	file, err := os.Open("./config/config.yaml")
	if err != nil {
		return err
	}

	decoder := yaml.NewDecoder(file)
	var conf config
	err = decoder.Decode(&conf)
	if err != nil {
		return err
	}
	conf.isConfigInit = true
	configInstance = conf

	return nil

}
func LoadAllConfig() error {
	err := read()
	return err
}

func GetDBConfig() DatabaseConfig {
	return configInstance.GameServer.Database
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

func IsEnableCachedHtml() bool {
	return configInstance.GameServer.Debug.EnabledCacheHTML
}

func IsEnableItems() bool {
	return configInstance.GameServer.Debug.EnabledItems
}
func IsEnableNPC() bool {
	return configInstance.GameServer.Debug.EnableNPC
}
func IsEnableSpawnList() bool {
	return configInstance.GameServer.Debug.EnabledSpawnlist
}
func IsEnableSkills() bool {
	return configInstance.GameServer.Debug.EnabledSkills
}
