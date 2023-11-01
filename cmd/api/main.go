package main

import (
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/server"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/pkg/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
)

func main() {
	log.Println("Starting api server")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	db, err := mysql.NewMysqlDB(viper.GetString("DB_DRIVER"), viper.GetString("DB_CONNECT"))

	s := server.NewServer(viper.GetString("APP_PORT"), db)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
