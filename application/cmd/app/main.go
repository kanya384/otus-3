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

	//db init
	pg, err := psql.New(cfg.DB.Host, cfg.DB.Port, cfg.DB.Name, cfg.DB.User, cfg.DB.Password, psql.MaxPoolSize(cfg.DB.PoolMax), psql.ConnTimeout(time.Duration(cfg.DB.Timeout)*time.Second))
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
