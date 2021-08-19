package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"l2gogameserver/config"
)

var db *pgxpool.Pool

func ConfigureDB() error {
	conf := config.Read()

	dsnString := "user=" + conf.LoginServer.Database.User
	dsnString += "password=" + conf.LoginServer.Database.Password
	dsnString += "host=" + conf.LoginServer.Database.Host
	dsnString += "port=" + conf.LoginServer.Database.Port
	dsnString += "dbname=" + conf.LoginServer.Database.Name
	dsnString += "sslmode=" + conf.LoginServer.Database.SSLMode
	dsnString += "pool_max_conns=" + conf.LoginServer.Database.PoolMaxConns

	pool, err := pgxpool.Connect(context.Background(), dsnString)
	if err != nil {
		return err
	}
	err = pool.Ping(context.Background())
	if err != nil {
		return err
	}
	db = pool
	return nil
}

func GetConn() (*pgxpool.Conn, error) {
	p, err := db.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	return p, nil
}
