package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var conn *sqlx.DB

func init() {
	db, err := sqlx.Connect("mysql", "root:root@(localhost:3306)/nordic")
	if err != nil {
		log.Fatalln("Could not connect to DB")
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	conn = db
}

func GetConn() *sqlx.DB {
	return conn
}

func SetConn(db *sqlx.DB) {
	conn = db
}
