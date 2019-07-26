package db

import "fmt"

type PostgresConfig struct {
	DBConfig
}

func (cfg *PostgresConfig) BuildConnString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Password,
		cfg.Username,
		cfg.DbName,
	)
}