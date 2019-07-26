package main

import (
	"basic/config"
	"basic/db"
	"basic/routes"
)

func main() {
	dbConfig := db.MysqlConfig{}
	dbConfig.Host = config.MYSQL_HOST
	dbConfig.Username = config.MYSQL_USER
	dbConfig.Password = config.MYSQL_PASS
	dbConfig.DbName = config.MYSQL_DB
	dbConfig.Port = "3306"
	dbConfig.MaxConn = 10

	db.Init("mysql", &dbConfig)
	routes.Init()
	routes.Start()
}