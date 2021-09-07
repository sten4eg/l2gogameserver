package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"l2gogameserver/config"
)

var db *pgxpool.Pool

func ConfigureDB() {
	conf := config.Read()

	dsnString := "user=" + conf.GameServer.Database.User
	dsnString += " password=" + conf.GameServer.Database.Password
	dsnString += " host=" + conf.GameServer.Database.Host
	dsnString += " port=" + conf.GameServer.Database.Port
	dsnString += " dbname=" + conf.GameServer.Database.Name
	dsnString += " sslmode=" + conf.GameServer.Database.SSLMode
	dsnString += " pool_max_conns=" + conf.GameServer.Database.PoolMaxConns

	pool, err := pgxpool.Connect(context.Background(), dsnString)
	if err != nil {
		panic(err)
	}
	err = pool.Ping(context.Background())
	if err != nil {
		panic(err)
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
