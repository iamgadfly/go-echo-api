package main

import (
	"github.com/AleksK1NG/api-mc/pkg/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iamgadfly/go-echo-api/config"
	"github.com/iamgadfly/go-echo-api/internal/server"
	mysql "github.com/iamgadfly/go-echo-api/pkg/mysql"
	"go.uber.org/zap"
	"log"
)

func main() {
	configPath := utils.GetConfigPath("")
	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	sugar := logger.Sugar()
	db, err := mysql.NewMysqlDB("mysql", cfg.Mysql.MysqlConnect)
	sugar.Info("starting server!")
	s := server.NewServer(cfg, cfg.Server.Port, db, sugar)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
