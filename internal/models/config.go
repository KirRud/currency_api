package models

type Config struct {
	DB     DBConfig
	Server ServerConfig
}

type DBConfig struct {
	DataBase string
}

type ServerConfig struct {
	Port string
}
