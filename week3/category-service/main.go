package main

import (
	"category-service/config"
	"category-service/dberr"
	"category-service/routes"
	"github.com/chuongdang/golang-libs/db"
	"os"
)

func main() {
	InitDB()
	go CheckDB()
	routes.Start()
}

func InitDB() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = config.Config.Mysql.Host
	}
	dbConfig := db.DBConfig{
		Username: config.Config.Mysql.Username,
		Password: config.Config.Mysql.Password,
		Host:     host,
		Port:     config.Config.Mysql.Port,
		DbName:   config.Config.Mysql.DbName,
		MaxConn:  config.Config.Mysql.MaxOpenConnection,
	}
	db.Init(&dbConfig)
}

func CheckDB() {
	for {
		err := <- dberr.ErrChannel
		if err == true {
			InitDB()
		}
	}
}

