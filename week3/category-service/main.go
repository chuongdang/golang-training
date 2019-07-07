package main

import (
	"category-service/config"
	"category-service/routes"

	"github.com/chuongdang/golang-libs/db"
)

func main() {
	dbConfig := db.DBConfig{
		Username: config.MYSQL_USERNAME,
		Password: config.MYSQL_PASSWORD,
		Host:     config.MYSQL_HOST,
		Port:     config.MYSQL_PORT,
		DbName:   config.MYSQL_DBNAME,
		MaxConn:  config.MYSQL_MAX_CONNECTION,
	}
	db.Init(&dbConfig)
	routes.Start()
}
