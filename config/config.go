package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	GameServer   GameServer `yaml:"gameserver"`
	isConfigInit bool
}
type GameServer struct {
	ServerIp           string         `yaml:"serverIp"`
	ServerId           int            `yaml:"serverId"`
	AcceptAlternateId  bool           `yaml:"acceptAlternateId"`
	ReserveHostOnLogin bool           `yaml:"reserveHostOnLogin"`
	Port               int            `yaml:"port"`
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

var configInstance Config

const MaxAdena = 99_900_000_000
const AdenaId = 57
const AncientAdenaId = 5575

const GeoFirstX = 11
const GeoFirstY = 10
const GeoLastX = 26
const GeoLastY = 26

//	Весь мир поделен на регионы, размер региона в клиенте равен размеру карты, а именно 32768x32768, диапазон Z от -32768 до 32767, идентификация карт в клиенте имеет вид XX_YY.
//
// Для более удобной работы с объектами на сервере, мир поделен на регионы, как по горизонтали так и по вертикали. Размер региона и ближайших его соседей соотвествует области видимости игрока.
// При настройке следует помнить: чем меньше размер региона, тем меньше нагрузка на процессор, тем меньше область видимости игрока, тем меньше исходящего трафика, но тем больше потребление памяти
// Данный параметр определяет размер региона по горизонтали: 1 << n,  при значении n = 15 - соответсвует размеру карты клиента,  при значении 12 размер равен 4096, 11 - 2048
const SHIFT_BY = 11

// Данный параметр определяет высоту региона по вертикали, при значении 10 - высота равна 1024
const SHIFT_BY_Z = 11

//func Get() GameServer {
//	if !configInstance.isConfigInit {
//		read()
//	}
//	return configInstance.GameServer
//}

func Read() (Config, error) {
	var conf Config
	file, err := os.Open("./config/config.yaml")
	if err != nil {
		return conf, err
	}

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(&conf)
	if err != nil {
		return conf, err
	}
	conf.isConfigInit = true
	configInstance = conf

	return conf, nil

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

func GetServerIp() string {
	return configInstance.GameServer.ServerIp
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

func GetPort() int {
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
