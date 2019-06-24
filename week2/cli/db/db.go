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
	db.DB.SetMaxOpenConns(10)   // The default is 0 (unlimited)
	db.DB.SetMaxIdleConns(10)   // defaultMaxIdleConns = 2
	db.DB.SetConnMaxLifetime(0) // 0, connections are reused forever.
	conn = db
}

func GetConn() *sqlx.DB {
	return conn
}

func SetConn(db *sqlx.DB) {
	conn = db
}
