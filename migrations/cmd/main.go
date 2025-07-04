package main

import (
	"embed"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"l2gogameserver/config"
	"l2gogameserver/migrations"
	"log"
	"net/url"
	"strconv"
)

func main() {

	err := config.Read()
	cfg := config.GetDBConfig()
	if err != nil {
		panic(err)
	}

	port, err := strconv.Atoi(cfg.Port)
	if err != nil {
		panic(err)
	}
	err = Migrate(Config{
		Credentials: Credentials{
			Username: cfg.User,
			Password: cfg.Password,
			Host:     cfg.Host,
			Port:     port,
			Database: cfg.Name,
		},
		Fs:     migrations.FS,
		FsPath: ".",
	})
	if err != nil {
		log.Fatal(err)
	}
}

type (
	// Credentials - data required to perform migration
	Credentials struct {
		Username string
		Password string
		Host     string
		Port     int
		Database string
		Scheme   string
	}
)

type (
	// Config configuration for migrations when reading from the embedded file system
	Config struct {
		Credentials
		Fs     embed.FS
		FsPath string
	}
)

// Migrate - performing migration
func Migrate(config Config) error {
	driver, err := iofs.New(config.Fs, config.FsPath)
	if err != nil {
		return err
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&search_path=%s",
		url.QueryEscape(config.Username), url.QueryEscape(config.Password), config.Host, config.Port, config.Database, config.Scheme)

	m, err := migrate.NewWithSourceInstance("iofs", driver, dsn)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
