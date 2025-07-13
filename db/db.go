package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"l2gogameserver/config"
	"l2gogameserver/data/logger"
	"strconv"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

var (
	ErrInvalidConfig      = errors.New("неверная конфигурация базы данных")
	ErrConnectionFailed   = errors.New("не удалось подключиться к базе данных")
	ErrPoolCreationFailed = errors.New("не удалось создать пул соединений")
	ErrPingFailed         = errors.New("не удалось выполнить ping базы данных")
)

// DatabaseManager управляет подключениями к базе данных
type DatabaseManager struct {
	pool    *pgxpool.Pool
	db      *sql.DB
	config  config.DatabaseConfig
	mu      sync.RWMutex
	metrics *DBMetrics
}

// DBMetrics содержит метрики базы данных
type DBMetrics struct {
	TotalConnections   int64
	ActiveConnections  int64
	IdleConnections    int64
	MaxConnections     int64
	ConnectionErrors   int64
	QueryErrors        int64
	LastConnectionTime time.Time
	mu                 sync.RWMutex
}

// NewDBMetrics создает новые метрики базы данных
func NewDBMetrics() *DBMetrics {
	return &DBMetrics{
		LastConnectionTime: time.Now(),
	}
}

// UpdateMetrics обновляет метрики
func (m *DBMetrics) UpdateMetrics(pool *pgxpool.Pool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	stats := pool.Stat()
	m.TotalConnections = int64(stats.TotalConns())
	m.ActiveConnections = int64(stats.AcquiredConns())
	m.IdleConnections = int64(stats.IdleConns())
	m.MaxConnections = int64(stats.MaxConns())
	m.LastConnectionTime = time.Now()
}

// GetMetrics возвращает копию метрик
func (m *DBMetrics) GetMetrics() DBMetrics {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return DBMetrics{
		TotalConnections:   m.TotalConnections,
		ActiveConnections:  m.ActiveConnections,
		IdleConnections:    m.IdleConnections,
		MaxConnections:     m.MaxConnections,
		ConnectionErrors:   m.ConnectionErrors,
		QueryErrors:        m.QueryErrors,
		LastConnectionTime: m.LastConnectionTime,
	}
}

// IncrementConnectionErrors увеличивает счетчик ошибок подключения
func (m *DBMetrics) IncrementConnectionErrors() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.ConnectionErrors++
}

// IncrementQueryErrors увеличивает счетчик ошибок запросов
func (m *DBMetrics) IncrementQueryErrors() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.QueryErrors++
}

// ConfigureDB создает подключение к базе данных
func ConfigureDB(config config.DatabaseConfig) (*sql.DB, error) {
	manager, err := NewDatabaseManager(config)
	if err != nil {
		return nil, err
	}

	return manager.GetDB(), nil
}

// NewDatabaseManager создает новый менеджер базы данных
func NewDatabaseManager(config config.DatabaseConfig) (*DatabaseManager, error) {
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidConfig, err)
	}

	manager := &DatabaseManager{
		config:  config,
		metrics: NewDBMetrics(),
	}

	if err := manager.connect(); err != nil {
		return nil, err
	}

	// Запускаем мониторинг метрик
	go manager.monitorMetrics()

	return manager, nil
}

// connect устанавливает подключение к базе данных
func (dm *DatabaseManager) connect() error {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	dsnString := dm.buildDSN()

	dbConfig, err := pgxpool.ParseConfig(dsnString)
	if err != nil {
		dm.metrics.IncrementConnectionErrors()
		return fmt.Errorf("%w: %v", ErrPoolCreationFailed, err)
	}

	// Настройки пула
	dbConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
	dbConfig.ConnConfig.ConnectTimeout = 10 * time.Second
	dbConfig.ConnConfig.RuntimeParams["application_name"] = "l2gogameserver"

	// Создаем пул
	pool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		dm.metrics.IncrementConnectionErrors()
		return fmt.Errorf("%w: %v", ErrConnectionFailed, err)
	}

	// Проверяем подключение
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		dm.metrics.IncrementConnectionErrors()
		return fmt.Errorf("%w: %v", ErrPingFailed, err)
	}

	dm.pool = pool
	dm.db = stdlib.OpenDBFromPool(pool)

	// Настраиваем максимальное количество соединений
	maxConns, err := strconv.Atoi(dm.config.PoolMaxConns)
	if err != nil {
		return fmt.Errorf("неверное значение pool_max_conns: %w", err)
	}

	dm.db.SetMaxOpenConns(maxConns)
	dm.db.SetMaxIdleConns(maxConns / 2)
	dm.db.SetConnMaxLifetime(30 * time.Minute)
	dm.db.SetConnMaxIdleTime(10 * time.Minute)

	logger.Info.Printf("Подключение к базе данных установлено: %s:%s/%s",
		dm.config.Host, dm.config.Port, dm.config.Name)

	return nil
}

// buildDSN строит строку подключения к базе данных
func (dm *DatabaseManager) buildDSN() string {
	dsnString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s search_path=%s pool_max_conns=%s",
		dm.config.User,
		dm.config.Password,
		dm.config.Host,
		dm.config.Port,
		dm.config.Name,
		dm.config.SSLMode,
		dm.config.Schema,
		dm.config.PoolMaxConns)

	return dsnString
}

// GetDB возвращает объект базы данных
func (dm *DatabaseManager) GetDB() *sql.DB {
	dm.mu.RLock()
	defer dm.mu.RUnlock()
	return dm.db
}

// GetPool возвращает пул соединений
func (dm *DatabaseManager) GetPool() *pgxpool.Pool {
	dm.mu.RLock()
	defer dm.mu.RUnlock()
	return dm.pool
}

// GetMetrics возвращает метрики базы данных
func (dm *DatabaseManager) GetMetrics() DBMetrics {
	return dm.metrics.GetMetrics()
}

// monitorMetrics мониторит метрики базы данных
func (dm *DatabaseManager) monitorMetrics() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		dm.mu.RLock()
		pool := dm.pool
		dm.mu.RUnlock()

		if pool != nil {
			dm.metrics.UpdateMetrics(pool)

			// Логируем метрики если есть проблемы
			metrics := dm.metrics.GetMetrics()
			if metrics.ActiveConnections > metrics.MaxConnections*80/100 {
				logger.Warning.Printf("Высокая нагрузка на БД: активных соединений %d из %d",
					metrics.ActiveConnections, metrics.MaxConnections)
			}
		}
	}
}

// Close закрывает все соединения
func (dm *DatabaseManager) Close() error {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	if dm.pool != nil {
		dm.pool.Close()
	}

	if dm.db != nil {
		return dm.db.Close()
	}

	return nil
}

// Transaction выполняет транзакцию
func (dm *DatabaseManager) Transaction(ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := dm.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("не удалось начать транзакцию: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("ошибка отката транзакции: %v, исходная ошибка: %w", rbErr, err)
		}
		return err
	}

	return tx.Commit()
}
