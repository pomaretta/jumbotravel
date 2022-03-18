package config

type Config struct {
	Database []DBConfig
}

type DBConfig struct {
	Host         string
	Port         string
	DatabaseName string
	User         string
	Password     string
	Enviroment   string
}
