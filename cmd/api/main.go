package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/iamgadfly/go-echo-api/config"
	"github.com/iamgadfly/go-echo-api/internal/grpcserver"
	"github.com/iamgadfly/go-echo-api/internal/server"
	"github.com/iamgadfly/go-echo-api/pkg/mysql"
	"go.uber.org/zap"
	"log"
	"time"
)

// @title App Parser
// @version 1.0
// @description app for parse
// @host localhost:8000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfgFile, err := config.LoadConfig("config/config-local.yml")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	if cfg.Mysql.IsDocker {
		time.Sleep(time.Second * 10)
	}

	db, err := mysql.NewMysqlDB("mysql", cfg.Mysql.MysqlConnect)
	if err != nil {
		log.Fatal(err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	sugar := logger.Sugar()
	er := make(chan error, 2)

	grpcServer := grpcserver.NewServerGRPC(db, cfg, sugar)
	s := server.NewServer(cfg, cfg.Server.Port, db, sugar)

	go grpcServer.Run(er)
	go s.Run(er)

	select {
	case err = <-er:
		log.Fatalf("Server error: %v\n", err.Error())
	}
}
