package db

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DbName   string
	MaxConn  int
}

type DBConfigInterface interface {
	BuildConnString() string
}
