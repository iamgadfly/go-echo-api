package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/iamgadfly/go-echo-api/config"
	"github.com/iamgadfly/go-echo-api/internal/server"
	mysql "github.com/iamgadfly/go-echo-api/pkg/mysql"
	"go.uber.org/zap"
	"log"
)

func main() {
	log.Println("Starting api server")
	cfgFile, err := config.LoadConfig("config/config.yml")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
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
	sugar.Info("starting server!")
	s := server.NewServer(cfg, cfg.Server.Port, db, sugar)

	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
