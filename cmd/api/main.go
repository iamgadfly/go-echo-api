package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/iamgadfly/go-echo-api/config"
	"github.com/iamgadfly/go-echo-api/internal/server"
	mysql "github.com/iamgadfly/go-echo-api/pkg/mysql"
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
	log.Println("Starting api server")
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

	logger, er := zap.NewProduction()
	if er != nil {
		log.Fatal(err)
	}

	sugar := logger.Sugar()
	s := server.NewServer(cfg, cfg.Server.Port, db, sugar)

	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
