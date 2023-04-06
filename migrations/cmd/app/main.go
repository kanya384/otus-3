package main

import (
	"fmt"
	"migrations/internal/config/config"

	"gitlab.com/kanya384/gotools/migrations"
)

func main() {
	cfg, err := config.InitConfig("")
	if err != nil {
		panic(fmt.Sprintf("error initializing config %s", err))
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)
	migrate, err := migrations.NewMigrations(dsn, "file://migrations")
	if err != nil {
		panic(fmt.Sprintf("migrations error: %s", err.Error()))

	}

	err = migrate.Up()
	if err != nil {
		panic(fmt.Sprintf("migrations error: %s", err.Error()))
	}
}
