package db

import "fmt"

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DbName   string
	MaxConn  int
}

func (cfg *DBConfig) BuildConnString() string {
	return fmt.Sprintf(
		"%v:%v@(%v:%v)/%v",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
}
