package db

import (
	"fmt"
	"log"
	"product-service/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var conn *sqlx.DB

func init() {
	connString := fmt.Sprintf(
		"%v:%v@(%v:%v)/%v",
		config.MYSQL_USERNAME,
		config.MYSQL_PASSWORD,
		config.MYSQL_HOST,
		config.MYSQL_PORT,
		config.MYSQL_DBNAME,
	)
	db, err := sqlx.Connect("mysql", connString)
	if err != nil {
		log.Fatalln("Could not connect to DB")
	}
	db.SetMaxIdleConns(config.MYSQL_MAX_CONNECTION)
	db.SetMaxOpenConns(config.MYSQL_MAX_CONNECTION)
	conn = db
}

func GetConn() *sqlx.DB {
	return conn
}

func SetConn(db *sqlx.DB) {
	conn = db
}
