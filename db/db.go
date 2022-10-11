package db

import (
	"context"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"l2gogameserver/config"
	"log"
	"os"
)

var db *pgxpool.Pool

func ConfigureDB() {
	conf := config.GetDBConfig()
	dsnString := "user=" + conf.User
	dsnString += " password=" + conf.Password
	dsnString += " host=" + conf.Host
	dsnString += " port=" + conf.Port
	dsnString += " dbname=" + conf.Name
	dsnString += " sslmode=" + conf.SSLMode
	dsnString += " search_path=" + conf.Schema
	dsnString += " pool_max_conns=" + conf.PoolMaxConns

	dbConfig, err := pgxpool.ParseConfig(dsnString)
	if err != nil {
		log.Fatal(err)
	}

	logrusLogger := &logrus.Logger{
		Out:          os.Stderr,
		Formatter:    new(logrus.JSONFormatter),
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.InfoLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}
	dbConfig.ConnConfig.Logger = logrusadapter.NewLogger(logrusLogger)
	conn, err := pgxpool.ConnectConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//pool, err := pgxpool.Connect(context.Background(), dsnString)
	//if err != nil {
	//	logger.Error.Panicln(err)
	//}
	//
	//err = pool.Ping(context.Background())
	//if err != nil {
	//	logger.Error.Panicln(err)
	//}
	db = conn
}

func GetConn() (*pgxpool.Conn, error) {
	p, err := db.Acquire(context.Background())
	if err != nil {
		return nil, err
	}

	return p, nil
}
