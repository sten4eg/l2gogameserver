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

const ENABLE_VIRTUAL_TERMINAL_PROCESSING = 0x0004

var (
	kernel32           = syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleMode = kernel32.NewProc("GetConsoleMode")
	procSetConsoleMode = kernel32.NewProc("SetConsoleMode")
	procGetStdHandle   = kernel32.NewProc("GetStdHandle")
)

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
	LevelClientToServer
	LevelServerToClient
)

const (
	colorReset      = "\033[0m"
	colorRed        = "\033[31m"
	colorGreen      = "\033[32m"
	colorYellow     = "\033[33m"
	colorBlue       = "\033[34m"
	colorPurple     = "\033[35m"
	colorCyan       = "\033[36m"
	colorWhite      = "\033[37m"
	colorLightRed   = "\033[91m"
	colorLightGreen = "\033[92m"
)

func enableWindowsColors() bool {
	if runtime.GOOS != "windows" {
		return false
	}
	handle, _, _ := procGetStdHandle.Call(0xfffffff5)
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

func isColorTerminal() bool {
	if runtime.GOOS == "windows" && enableWindowsColors() {
		return true
	}
	t := os.Getenv("TERM")
	if t == "" {
		return false
	}
	fi, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}

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

type Config struct {
	Level      Level  `json:"level"`
	OutputPath string `json:"output_path"`
	MaxSize    int64  `json:"max_size"`
	MaxBackups int    `json:"max_backups"`
	MaxAge     int    `json:"max_age"`
	Compress   bool   `json:"compress"`
	UseColors  bool   `json:"use_colors"`
}

type Logger struct {
	config Config
	mu     sync.RWMutex
	output io.Writer
	level  Level
}

var (
	Warning      *log.Logger
	Info         *log.Logger
	Error        *log.Logger
	globalLogger *Logger
	once         sync.Once
)

func GetLogger() *Logger {
	once.Do(func() {
		useColors := isColorTerminal()
		globalLogger = NewLogger(Config{
			Level:      LevelInfo,
			OutputPath: "",
			MaxSize:    100 * 1024 * 1024,
			MaxBackups: 5,
			MaxAge:     30,
			Compress:   true,
			UseColors:  useColors,
		})
	})
	return globalLogger
}

func NewLogger(config Config) *Logger {
	logger := &Logger{
		config: config,
		level:  config.Level,
	}
	if config.OutputPath != "" {
		writer, err := logger.createRotatingWriter()
		if err != nil {
			logger.output = os.Stdout
		} else {
			logger.output = writer
		}
	} else {
		logger.output = os.Stdout
	}
	return logger
}

func (l *Logger) createRotatingWriter() (io.Writer, error) {
	dir := filepath.Dir(l.config.OutputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("не удалось создать директорию: %w", err)
	}
	file, err := os.OpenFile(l.config.OutputPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл логов: %w", err)
	}
	return file, nil
}

func (l *Logger) Log(level Level, format string, args ...interface{}) {
	if level < l.level {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()

	// ВАЖНО: уровень Caller = 4, чтобы показать откуда вызван LogInfo
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "unknown"
		line = 0
	}
	parts := strings.Split(file, "/")
	fileName := parts[len(parts)-1]
	timestamp := time.Now().Format("15:04:05")
	message := fmt.Sprintf(format, args...)
	logMessage := fmt.Sprintf("%s %s:%d: %s", timestamp, fileName, line, message)

	if l.config.UseColors && isColorTerminal() {
		logMessage = level.getColor() + logMessage + colorReset
	}

	fmt.Fprintln(l.output, logMessage)
}

func (l *Logger) Debug(format string, args ...interface{})   { l.Log(LevelDebug, format, args...) }
func (l *Logger) Info(format string, args ...interface{})    { l.Log(LevelInfo, format, args...) }
func (l *Logger) Warning(format string, args ...interface{}) { l.Log(LevelWarning, format, args...) }
func (l *Logger) Error(format string, args ...interface{})   { l.Log(LevelError, format, args...) }
func (l *Logger) Fatal(format string, args ...interface{}) {
	l.Log(LevelFatal, format, args...)
	os.Exit(1)
}
func (l *Logger) ClientToServer(format string, args ...interface{}) {
	l.Log(LevelClientToServer, format, args...)
}
func (l *Logger) ServerToClient(format string, args ...interface{}) {
	l.Log(LevelServerToClient, format, args...)
}

func (l *Logger) SetLevel(level Level)  { l.mu.Lock(); defer l.mu.Unlock(); l.level = level }
func (l *Logger) GetLevel() Level       { l.mu.RLock(); defer l.mu.RUnlock(); return l.level }
func (l *Logger) SetUseColors(use bool) { l.mu.Lock(); defer l.mu.Unlock(); l.config.UseColors = use }
func (l *Logger) GetUseColors() bool    { l.mu.RLock(); defer l.mu.RUnlock(); return l.config.UseColors }
func (l *Logger) ForceEnableColors()    { l.mu.Lock(); defer l.mu.Unlock(); l.config.UseColors = true }

type ContextLogger struct {
	logger *Logger
	ctx    context.Context
}

func (cl *ContextLogger) getRequestID() string {
	if cl.ctx == nil {
		return ""
	}
	if requestID, ok := cl.ctx.Value("request_id").(string); ok {
		return requestID
	}
	return ""
}

func (cl *ContextLogger) Log(level Level, format string, args ...interface{}) {
	if requestID := cl.getRequestID(); requestID != "" {
		format = "[RequestID: %s] " + format
		args = append([]interface{}{requestID}, args...)
	}
	cl.logger.Log(level, format, args...)
}

func (cl *ContextLogger) Debug(format string, args ...interface{}) {
	cl.Log(LevelDebug, format, args...)
}
func (cl *ContextLogger) Info(format string, args ...interface{}) { cl.Log(LevelInfo, format, args...) }
func (cl *ContextLogger) Warning(format string, args ...interface{}) {
	cl.Log(LevelWarning, format, args...)
}
func (cl *ContextLogger) Error(format string, args ...interface{}) {
	cl.Log(LevelError, format, args...)
}
func (cl *ContextLogger) Fatal(format string, args ...interface{}) {
	cl.Log(LevelFatal, format, args...)
	os.Exit(1)
}
func (cl *ContextLogger) ClientToServer(format string, args ...interface{}) {
	cl.Log(LevelClientToServer, format, args...)
}
func (cl *ContextLogger) ServerToClient(format string, args ...interface{}) {
	cl.Log(LevelServerToClient, format, args...)
}

func (l *Logger) WithContext(ctx context.Context) *ContextLogger {
	return &ContextLogger{logger: l, ctx: ctx}
}

func init() {
	logger := GetLogger()
	Info = log.New(logger.output, "INFO: ", log.Ltime|log.Lshortfile)
	Warning = log.New(logger.output, "WARNING: ", log.Ltime|log.Lshortfile)
	Error = log.New(logger.output, "ERROR: ", log.Ltime|log.Lshortfile)
	logger.ForceEnableColors()
	logger.Info("Логгер инициализирован")
}

func LogDebug(format string, args ...interface{})   { GetLogger().Debug(format, args...) }
func LogInfo(format string, args ...interface{})    { GetLogger().Info(format, args...) }
func LogWarning(format string, args ...interface{}) { GetLogger().Warning(format, args...) }
func LogError(format string, args ...interface{})   { GetLogger().Error(format, args...) }
func LogFatal(format string, args ...interface{})   { GetLogger().Fatal(format, args...) }
func LogClientToServer(format string, args ...interface{}) {
	GetLogger().ClientToServer(format, args...)
}
func LogServerToClient(format string, args ...interface{}) {
	GetLogger().ServerToClient(format, args...)
}
