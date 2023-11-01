package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

func NewMysqlDB(driver string, conn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(driver, conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db, nil

}
