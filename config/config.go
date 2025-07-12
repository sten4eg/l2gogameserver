package config

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

// Константы для конфигурации
const (
	MaxAdena       = 99_900_000_000
	AdenaId        = 57
	AncientAdenaId = 5575

	GeoFirstX = 11
	GeoFirstY = 10
	GeoLastX  = 26
	GeoLastY  = 26

	// Весь мир поделен на регионы, размер региона в клиенте равен размеру карты, а именно 32768x32768,
	// диапазон Z от -32768 до 32767, идентификация карт в клиенте имеет вид XX_YY.
	//
	// Для более удобной работы с объектами на сервере, мир поделен на регионы, как по горизонтали так и по вертикали.
	// Размер региона и ближайших его соседей соответствует области видимости игрока.
	// При настройке следует помнить: чем меньше размер региона, тем меньше нагрузка на процессор,
	// тем меньше область видимости игрока, тем меньше исходящего трафика, но тем больше потребление памяти
	// Данный параметр определяет размер региона по горизонтали: 1 << n,
	// при значении n = 15 - соответствует размеру карты клиента,
	// при значении 12 размер равен 4096, 11 - 2048
	SHIFT_BY = 11

	// Данный параметр определяет высоту региона по вертикали, при значении 10 - высота равна 1024
	SHIFT_BY_Z = 11

	// Пути к файлам конфигурации
	DefaultConfigPath = "./config/config.yaml"
)

var (
	ErrConfigNotInitialized = errors.New("конфигурация не инициализирована")
	ErrInvalidConfigPath    = errors.New("неверный путь к файлу конфигурации")
	ErrConfigFileNotFound   = errors.New("файл конфигурации не найден")
	ErrInvalidConfigData    = errors.New("неверные данные конфигурации")
)

// Config представляет основную конфигурацию приложения
type Config struct {
	GameServer   GameServer `yaml:"gameserver"`
	isConfigInit bool
	mu           sync.RWMutex
}

// GameServer представляет конфигурацию игрового сервера
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

// DatabaseConfig представляет конфигурацию базы данных
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

// Debug представляет настройки отладки
type Debug struct {
	ShowPackets      bool   `yaml:"show_packets"`
	IgnorePackets    []byte `yaml:"ignore_packets"`
	EnableNPC        bool   `yaml:"enable_load_npc"`
	EnabledSkills    bool   `yaml:"enabled_load_skills"`
	EnabledItems     bool   `yaml:"enabled_items"`
	EnabledSpawnlist bool   `yaml:"enabled_spawnlist"`
	// EnabledCacheHTML - если false, тогда не будет записываться в кэш, удобно для
	// редактирования HTML диалогов и просмотра в игре при каждом обращении.
	EnabledCacheHTML bool `yaml:"enabled_cache_html"`
}

// ConfigManager управляет конфигурацией приложения
type ConfigManager struct {
	config Config
	mu     sync.RWMutex
}

var (
	configInstance *ConfigManager
	once           sync.Once
)

// GetConfigManager возвращает синглтон менеджера конфигурации
func GetConfigManager() *ConfigManager {
	once.Do(func() {
		configInstance = &ConfigManager{}
	})
	return configInstance
}

// Read читает конфигурацию из файла
func Read() (Config, error) {
	return ReadFromFile(DefaultConfigPath)
}

// ReadFromFile читает конфигурацию из указанного файла
func ReadFromFile(configPath string) (Config, error) {
	var conf Config

	if configPath == "" {
		return conf, ErrInvalidConfigPath
	}

	// Проверяем существование файла
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return conf, fmt.Errorf("%w: %s", ErrConfigFileNotFound, configPath)
	}

	// Открываем файл
	file, err := os.Open(configPath)
	if err != nil {
		return conf, fmt.Errorf("ошибка открытия файла конфигурации: %w", err)
	}
	defer file.Close()

	// Декодируем YAML
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&conf); err != nil {
		return conf, fmt.Errorf("%w: %v", ErrInvalidConfigData, err)
	}

	// Валидируем конфигурацию
	if err := conf.Validate(); err != nil {
		return conf, fmt.Errorf("ошибка валидации конфигурации: %w", err)
	}

	conf.isConfigInit = true
	return conf, nil
}

// Validate валидирует конфигурацию
func (c *Config) Validate() error {
	if c.GameServer.ServerIp == "" {
		return errors.New("ServerIp не может быть пустым")
	}

	if c.GameServer.Port <= 0 || c.GameServer.Port > 65535 {
		return errors.New("Port должен быть в диапазоне 1-65535")
	}

	if c.GameServer.ServerId < 0 {
		return errors.New("ServerId не может быть отрицательным")
	}

	if c.GameServer.MaxPlayer <= 0 {
		return errors.New("MaxPlayer должен быть больше 0")
	}

	// Валидация конфигурации базы данных
	if err := c.GameServer.Database.Validate(); err != nil {
		return fmt.Errorf("ошибка конфигурации базы данных: %w", err)
	}

	return nil
}

// Validate валидирует конфигурацию базы данных
func (d *DatabaseConfig) Validate() error {
	if d.Name == "" {
		return errors.New("имя базы данных не может быть пустым")
	}

	if d.Host == "" {
		return errors.New("хост базы данных не может быть пустым")
	}

	if d.Port == "" {
		return errors.New("порт базы данных не может быть пустым")
	}

	if d.User == "" {
		return errors.New("пользователь базы данных не может быть пустым")
	}

	return nil
}

// GetConnectionString возвращает строку подключения к базе данных
func (d *DatabaseConfig) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSLMode)
}

// Load загружает конфигурацию в менеджер
func (cm *ConfigManager) Load(configPath string) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	config, err := ReadFromFile(configPath)
	if err != nil {
		return err
	}

	cm.config = config
	return nil
}

// Get возвращает текущую конфигурацию
func (cm *ConfigManager) Get() Config {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	if !cm.config.isConfigInit {
		panic(ErrConfigNotInitialized)
	}

	return cm.config
}

// GetGameServer возвращает конфигурацию игрового сервера
func (cm *ConfigManager) GetGameServer() GameServer {
	return cm.Get().GameServer
}

// GetDatabase возвращает конфигурацию базы данных
func (cm *ConfigManager) GetDatabase() DatabaseConfig {
	return cm.GetGameServer().Database
}

// GetDebug возвращает настройки отладки
func (cm *ConfigManager) GetDebug() Debug {
	return cm.GetGameServer().Debug
}

// Reload перезагружает конфигурацию из файла
func (cm *ConfigManager) Reload(configPath string) error {
	return cm.Load(configPath)
}

// Устаревшие функции для обратной совместимости
// Deprecated: используйте ConfigManager

func GetDBConfig() DatabaseConfig {
	return GetConfigManager().GetDatabase()
}

func GetHexId() []byte {
	return GetConfigManager().GetGameServer().HexId
}

func GetLoginServerPort() string {
	return GetConfigManager().GetGameServer().PortForLS
}

func GetServerIp() string {
	return GetConfigManager().GetGameServer().ServerIp
}

func GetServerId() int {
	return GetConfigManager().GetGameServer().ServerId
}

func GetAcceptAlternateId() bool {
	return GetConfigManager().GetGameServer().AcceptAlternateId
}

func GetReserveHostOnLogin() bool {
	return GetConfigManager().GetGameServer().ReserveHostOnLogin
}

func GetPort() int {
	return GetConfigManager().GetGameServer().Port
}

func GetMaxPlayer() int {
	return GetConfigManager().GetGameServer().MaxPlayer
}

func GetServerListBrackets() bool {
	return GetConfigManager().GetGameServer().ServerListBrackets
}

func GetGMOnly() bool {
	return GetConfigManager().GetGameServer().GMOnly
}

func GetServerListAge() byte {
	return GetConfigManager().GetGameServer().ServerListAge
}

func GetServerListType() string {
	return GetConfigManager().GetGameServer().ServerListType
}

func IsEnableCachedHtml() bool {
	return GetConfigManager().GetDebug().EnabledCacheHTML
}

func IsEnableItems() bool {
	return GetConfigManager().GetDebug().EnabledItems
}

func IsEnableNPC() bool {
	return GetConfigManager().GetDebug().EnableNPC
}

func IsEnableSpawnList() bool {
	return GetConfigManager().GetDebug().EnabledSpawnlist
}

func IsEnableSkills() bool {
	return GetConfigManager().GetDebug().EnabledSkills
}

func GetDebug() Debug {
	return GetConfigManager().GetDebug()
}

// IsShowPackets возвращает true если включен показ пакетов
func (d Debug) IsShowPackets() bool {
	return d.ShowPackets
}

// IsIgnorePackets возвращает true если пакет должен быть проигнорирован
func (d Debug) IsIgnorePackets(opcode byte) bool {
	// Здесь можно добавить логику для игнорирования определенных опкодов
	// Например, игнорировать пакеты с опкодами 0x00, 0xFF и т.д.
	for _, ignored := range d.IgnorePackets {
		if opcode == ignored {
			return true
		}
	}
	return false
}
