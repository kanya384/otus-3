package main

import (
	"fmt"
	"os"
	"os/signal"
	"otus/internal/config"
	"otus/internal/delivery"
	"otus/internal/repository"
	"otus/internal/service"
	"syscall"
	"time"

	lg "gitlab.com/kanya384/gotools/logger"
	"gitlab.com/kanya384/gotools/migrations"
	"gitlab.com/kanya384/gotools/psql"
)

func main() {
	cfg, err := config.InitConfig("")
	if err != nil {
		panic(fmt.Sprintf("error initializing config %s", err))
	}

	//setup logger
	logger, err := lg.New(cfg.Log.Level, cfg.App.ServiceName)
	if err != nil {
		panic(fmt.Sprintf("error initializing logger %s", err))
	}
	logFile, err := os.OpenFile(cfg.Log.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(fmt.Sprintf("error opening log file %s", err.Error()))
	}
	logger.SetOutput(logFile)

	//db migrations
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.PG.User, cfg.PG.Pass, cfg.PG.Host, cfg.PG.Port, cfg.PG.DbName)
	migrate, err := migrations.NewMigrations(dsn, "file://migrations")
	if err != nil {
		logger.Fatalf("migrations error: %s", err.Error())
	}

	err = migrate.Up()
	if err != nil {
		logger.Fatalf("migrations error: %s", err.Error())
	}

	//db init
	pg, err := psql.New(cfg.PG.Host, cfg.PG.Port, cfg.PG.DbName, cfg.PG.User, cfg.PG.Pass, psql.MaxPoolSize(cfg.PG.PoolMax), psql.ConnTimeout(time.Duration(cfg.PG.Timeout)*time.Second))
	if err != nil {
		logger.Fatal(fmt.Errorf("postgres connection error: %w", err))
	}

	//repository
	repository, err := repository.New(pg, repository.Options{})
	if err != nil {
		logger.Fatal("storage initialization error: %s", err.Error())
	}

	//service
	service, err := service.New(repository, logger, service.Options{})
	if err != nil {
		logger.Fatal("services initialization error: %s", err.Error())
	}

	delivery, err := delivery.New(service, delivery.Options{})
	if err != nil {
		logger.Fatal("delivery initialization error: %s", err.Error())
	}

	err = delivery.Run(cfg.App.Port)
	if err != nil {
		logger.Fatal("start delivery error: %s", err.Error())
	}

	//closes connections on app kill
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	if err := shutdown(pg, logger); err != nil {
		logger.Fatal(fmt.Errorf("failed shutdown with error: %w", err))
	}

}

func shutdown(psql *psql.Postgres, logger *lg.Logger) error {
	fmt.Println("Gracefull shut down in progress...")
	psql.Pool.Close()
	logger.Info("Gracefull shutdown done!")
	return nil
}