package main

import (
	"log"
	"os"
	"product-service/config"
	"product-service/routes"
	"time"

	"github.com/chuongdang/golang-libs/db"
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
		err := db.GetConn().Ping()
		if err != nil {
			log.Println(err)
			InitDB()
		}
		time.Sleep(1 * time.Second)
	}
}
