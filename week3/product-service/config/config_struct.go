package config

var (
	Config *AppConfig
)

type AppConfig struct {
	Api                 ApiConfig
	Mysql               MysqlConfig
	Services            ServicesConfig
}

type ApiConfig struct {
	Port string
}

type MysqlConfig struct {
	Host              string
	Port              string
	Username          string
	Password          string
	DbName            string
	Charset           string
	MaxIdleConnection int
	MaxOpenConnection int
}

type ServicesConfig struct {
	CategoryService string
}
