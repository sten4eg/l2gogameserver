package db

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"l2gogameserver/config"
	"strconv"
)

func ConfigureDB(config config.DatabaseConfig) (*sql.DB, error) {

	dsnString := "user=" + config.User
	dsnString += " password=" + config.Password
	dsnString += " host=" + config.Host
	dsnString += " port=" + config.Port
	dsnString += " dbname=" + config.Name
	dsnString += " sslmode=" + config.SSLMode
	dsnString += " search_path=" + config.Schema
	dsnString += " pool_max_conns=" + config.PoolMaxConns

	// unixWayPostgres := "postgresql:///postgres?host=/run/postgresql&port=5432&user=postgres&password=postgres&sslmode=disable"
	dbConfig, err := pgxpool.ParseConfig(dsnString)
	if err != nil {
		return nil, err
	}

	dbConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	pool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	db := stdlib.OpenDBFromPool(pool)
	maxConni, err := strconv.Atoi(config.PoolMaxConns)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxConni)
	return db, nil

}
