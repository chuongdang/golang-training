package db

import "fmt"

type MysqlConfig struct {
	DBConfig
}

func (cfg *MysqlConfig) BuildConnString() string {
	return fmt.Sprintf(
		"%v:%v@(%v:%v)/%v",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
}