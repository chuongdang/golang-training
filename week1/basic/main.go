package main

import (
	"basic/config"
	"basic/db"
	"basic/routes"
)

func main() {
	dbConfig := db.DBConfig{
		Host: config.MYSQL_HOST,
		Username: config.MYSQL_USER,
		Password: config.MYSQL_PASS,
		DbName: config.MYSQL_DB,
		Port: "3306",
		MaxConn: 10,
	}
	db.Init(&dbConfig)
	routes.Init()
	routes.Start()
}