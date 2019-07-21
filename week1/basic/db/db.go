package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var conn *sqlx.DB

func Init(config *DBConfig) {
	db, err := sqlx.Connect("mysql", config.BuildConnString())
	if err != nil {
		log.Fatalln("Could not connect to DB")
	}
	db.SetMaxIdleConns(config.MaxConn)
	db.SetMaxOpenConns(config.MaxConn)
	conn = db
}

func GetConn() *sqlx.DB {
	return conn
}

func SetConn(db *sqlx.DB) {
	conn = db
}