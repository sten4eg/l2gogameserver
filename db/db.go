package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"l2gogameserver/config"
	"l2gogameserver/data/logger"
)

var db *pgxpool.Pool

//Connect to DB (postgres)
func ConfigureDB() {
	conf := config.Get().Database
	dsnString := "user=" + conf.User
	dsnString += " password=" + conf.Password
	dsnString += " host=" + conf.Host
	dsnString += " port=" + conf.Port
	dsnString += " dbname=" + conf.Name
	dsnString += " sslmode=" + conf.SSLMode
	dsnString += " pool_max_conns=" + conf.PoolMaxConns

	pool, err := pgxpool.Connect(context.Background(), dsnString)
	if err != nil {
		logger.Error.Panicln(err)
	}
	err = pool.Ping(context.Background())
	if err != nil {
		logger.Error.Panicln(err)
	}
	db = pool
}

func GetConn() (*pgxpool.Conn, error) {
	p, err := db.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	return p, nil
}
