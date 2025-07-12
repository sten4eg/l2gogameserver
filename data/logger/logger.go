package logger

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"
	"unsafe"
)

// Windows API константы и структуры
const (
	ENABLE_VIRTUAL_TERMINAL_PROCESSING = 0x0004
)

var (
	kernel32           = syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleMode = kernel32.NewProc("GetConsoleMode")
	procSetConsoleMode = kernel32.NewProc("SetConsoleMode")
	procGetStdHandle   = kernel32.NewProc("GetStdHandle")
)

// Level представляет уровень логирования
type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
	LevelClientToServer // новый уровень для пакетов Client->Server
	LevelServerToClient // новый уровень для пакетов Server->Client
)

// Цветовые коды ANSI
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"

	// Специальные цвета для пакетов
	colorLightRed   = "\033[91m" // светло-красный для Client->Server
	colorLightGreen = "\033[92m" // светло-зеленый для Server->Client
)

// enableWindowsColors включает поддержку ANSI цветов в Windows
func enableWindowsColors() bool {
	if runtime.GOOS != "windows" {
		return false
	}

	handle, _, _ := procGetStdHandle.Call(0xfffffff5) // STD_OUTPUT_HANDLE = -11
	if handle == 0 {
		return false
	}

	var mode uint32
	ret, _, _ := procGetConsoleMode.Call(handle, uintptr(unsafe.Pointer(&mode)))
	if ret == 0 {
		return false
	}

	mode |= ENABLE_VIRTUAL_TERMINAL_PROCESSING
	ret, _, _ = procSetConsoleMode.Call(handle, uintptr(mode))
	return ret != 0
}

// isColorTerminal проверяет, поддерживает ли терминал цвета
func isColorTerminal() bool {
	// Для Windows пытаемся включить поддержку цветов
	if runtime.GOOS == "windows" {
		if enableWindowsColors() {
			return true
		}
	}

	// Проверяем переменную окружения TERM
	term := os.Getenv("TERM")
	if term == "" {
		return false
	}

	// Проверяем, что это интерактивный терминал
	fileInfo, err := os.Stdout.Stat()
	if err != nil {
		return false
	}

	// Проверяем, что stdout подключен к терминалу
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

// String возвращает строковое представление уровня логирования
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarning:
		return "WARNING"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	case LevelClientToServer:
		return "CLIENT->SERVER"
	case LevelServerToClient:
		return "SERVER->CLIENT"
	default:
		return "UNKNOWN"
	}
}

// getColor возвращает цвет для уровня логирования
func (l Level) getColor() string {
	switch l {
	case LevelDebug:
		return colorCyan
	case LevelInfo:
		return colorGreen
	case LevelWarning:
		return colorYellow
	case LevelError:
		return colorRed
	case LevelFatal:
		return colorPurple
	case LevelClientToServer:
		return colorLightRed
	case LevelServerToClient:
		return colorLightGreen
	default:
		return colorWhite
	}
}

// Config представляет конфигурацию логгера
type Config struct {
	Level      Level  `json:"level"`
	OutputPath string `json:"output_path"`
	MaxSize    int64  `json:"max_size"`    // максимальный размер файла в байтах
	MaxBackups int    `json:"max_backups"` // максимальное количество резервных файлов
	MaxAge     int    `json:"max_age"`     // максимальный возраст файла в днях
	Compress   bool   `json:"compress"`    // сжимать ли старые файлы
	UseColors  bool   `json:"use_colors"`  // использовать ли цвета в выводе
}

// Logger представляет улучшенный логгер
type Logger struct {
	config Config
	mu     sync.RWMutex
	output io.Writer
	level  Level
}

var (
	// Глобальные экземпляры логгеров для обратной совместимости
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger

	// Глобальный логгер
	globalLogger *Logger
	once         sync.Once
)

// GetLogger возвращает глобальный экземпляр логгера
func GetLogger() *Logger {
	once.Do(func() {
		if globalLogger == nil {
			// Проверяем поддержку цветов в терминале
			useColors := isColorTerminal()

			globalLogger = NewLogger(Config{
				Level:      LevelInfo,
				OutputPath: "",
				MaxSize:    100 * 1024 * 1024, // 100MB
				MaxBackups: 5,
				MaxAge:     30,
				Compress:   true,
				UseColors:  useColors, // автоматически определяем поддержку цветов
			})
		}
	})
	return globalLogger
}

// NewLogger создает новый логгер с указанной конфигурацией
func NewLogger(config Config) *Logger {
	logger := &Logger{
		config: config,
		level:  config.Level,
	}

	// Настраиваем вывод
	if config.OutputPath != "" {
		writer, err := logger.createRotatingWriter()
		if err != nil {
			// Если не удалось создать файл, используем stdout
			logger.output = os.Stdout
		} else {
			logger.output = writer
		}
	} else {
		logger.output = os.Stdout
	}

	return logger
}

// createRotatingWriter создает writer с ротацией файлов
func (l *Logger) createRotatingWriter() (io.Writer, error) {
	// Создаем директорию если не существует
	dir := filepath.Dir(l.config.OutputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("не удалось создать директорию для логов: %w", err)
	}

	// Открываем файл для записи
	file, err := os.OpenFile(l.config.OutputPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл логов: %w", err)
	}

	return file, nil
}

// Log записывает сообщение с указанным уровнем
func (l *Logger) Log(level Level, format string, args ...interface{}) {
	if level < l.level {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	// Получаем информацию о вызывающем коде
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "unknown"
		line = 0
	}

	// Формируем сообщение
	message := fmt.Sprintf(format, args...)
	timestamp := time.Now().Format("15:04:05")

	// Сокращаем путь к файлу - берем только имя файла
	parts := strings.Split(file, "/")
	fileName := parts[len(parts)-1]

	// Формируем сообщение в новом формате
	logMessage := fmt.Sprintf("%s %s:%d: %s",
		timestamp, fileName, line, message)

	// Добавляем цвета если включены и терминал поддерживает цвета
	if l.config.UseColors && isColorTerminal() {
		color := level.getColor()
		logMessage = color + logMessage + colorReset
	}

	fmt.Fprintln(l.output, logMessage)
}

// Debug записывает отладочное сообщение
func (l *Logger) Debug(format string, args ...interface{}) {
	l.Log(LevelDebug, format, args...)
}

// Info записывает информационное сообщение
func (l *Logger) Info(format string, args ...interface{}) {
	l.Log(LevelInfo, format, args...)
}

// Warning записывает предупреждение
func (l *Logger) Warning(format string, args ...interface{}) {
	l.Log(LevelWarning, format, args...)
}

// Error записывает ошибку
func (l *Logger) Error(format string, args ...interface{}) {
	l.Log(LevelError, format, args...)
}

// Fatal записывает критическую ошибку и завершает программу
func (l *Logger) Fatal(format string, args ...interface{}) {
	l.Log(LevelFatal, format, args...)
	os.Exit(1)
}

// ClientToServer записывает сообщение о пакете Client->Server
func (l *Logger) ClientToServer(format string, args ...interface{}) {
	l.Log(LevelClientToServer, format, args...)
}

// ServerToClient записывает сообщение о пакете Server->Client
func (l *Logger) ServerToClient(format string, args ...interface{}) {
	l.Log(LevelServerToClient, format, args...)
}

// WithContext создает логгер с контекстом
func (l *Logger) WithContext(ctx context.Context) *ContextLogger {
	return &ContextLogger{
		logger: l,
		ctx:    ctx,
	}
}

// SetLevel устанавливает уровень логирования
func (l *Logger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

// GetLevel возвращает текущий уровень логирования
func (l *Logger) GetLevel() Level {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.level
}

// SetUseColors устанавливает использование цветов
func (l *Logger) SetUseColors(useColors bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.config.UseColors = useColors
}

// GetUseColors возвращает настройку использования цветов
func (l *Logger) GetUseColors() bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.config.UseColors
}

// ForceEnableColors принудительно включает цвета (для отладки)
func (l *Logger) ForceEnableColors() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.config.UseColors = true
}

// ContextLogger представляет логгер с контекстом
type ContextLogger struct {
	logger *Logger
	ctx    context.Context
}

// Log записывает сообщение с контекстом
func (cl *ContextLogger) Log(level Level, format string, args ...interface{}) {
	// Добавляем информацию о контексте
	if requestID := cl.getRequestID(); requestID != "" {
		format = "[RequestID: %s] " + format
		args = append([]interface{}{requestID}, args...)
	}

	cl.logger.Log(level, format, args...)
}

// getRequestID извлекает ID запроса из контекста
func (cl *ContextLogger) getRequestID() string {
	if cl.ctx == nil {
		return ""
	}

	if requestID, ok := cl.ctx.Value("request_id").(string); ok {
		return requestID
	}

	return ""
}

// Debug записывает отладочное сообщение с контекстом
func (cl *ContextLogger) Debug(format string, args ...interface{}) {
	cl.Log(LevelDebug, format, args...)
}

// Info записывает информационное сообщение с контекстом
func (cl *ContextLogger) Info(format string, args ...interface{}) {
	cl.Log(LevelInfo, format, args...)
}

// Warning записывает предупреждение с контекстом
func (cl *ContextLogger) Warning(format string, args ...interface{}) {
	cl.Log(LevelWarning, format, args...)
}

// Error записывает ошибку с контекстом
func (cl *ContextLogger) Error(format string, args ...interface{}) {
	cl.Log(LevelError, format, args...)
}

// Fatal записывает критическую ошибку с контекстом и завершает программу
func (cl *ContextLogger) Fatal(format string, args ...interface{}) {
	cl.Log(LevelFatal, format, args...)
	os.Exit(1)
}

// ClientToServer записывает сообщение о пакете Client->Server с контекстом
func (cl *ContextLogger) ClientToServer(format string, args ...interface{}) {
	cl.Log(LevelClientToServer, format, args...)
}

// ServerToClient записывает сообщение о пакете Server->Client с контекстом
func (cl *ContextLogger) ServerToClient(format string, args ...interface{}) {
	cl.Log(LevelServerToClient, format, args...)
}

// Инициализация для обратной совместимости
func init() {
	logger := GetLogger()

	// Создаем стандартные логгеры для обратной совместимости
	Info = log.New(logger.output, "INFO: ", log.Ltime|log.Lshortfile)
	Warning = log.New(logger.output, "WARNING: ", log.Ltime|log.Lshortfile)
	Error = log.New(logger.output, "ERROR: ", log.Ltime|log.Lshortfile)

	// Принудительно включаем цвета для тестирования
	logger.ForceEnableColors()

	// Тестовое сообщение для проверки инициализации логгера
	logger.Info("Логгер инициализирован")
}

// Глобальные функции для удобства использования
func LogDebug(format string, args ...interface{}) {
	GetLogger().Debug(format, args...)
}

func LogInfo(format string, args ...interface{}) {
	GetLogger().Info(format, args...)
}

func LogWarning(format string, args ...interface{}) {
	GetLogger().Warning(format, args...)
}

func LogError(format string, args ...interface{}) {
	GetLogger().Error(format, args...)
}

func LogFatal(format string, args ...interface{}) {
	GetLogger().Fatal(format, args...)
}

// Новые глобальные функции для пакетов
func LogClientToServer(format string, args ...interface{}) {
	GetLogger().ClientToServer(format, args...)
}

func LogServerToClient(format string, args ...interface{}) {
	GetLogger().ServerToClient(format, args...)
}
