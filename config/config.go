package config

import (
	"currency_api/internal/models"
	"github.com/joho/godotenv"
	"os"
)

func InitConfig() (*models.Config, error) {
	var cfg models.Config
	err := godotenv.Load("./config/dev.env")
	if err != nil {
		return nil, err
	}

	cfg.DB = models.DBConfig{
		DataBase: os.Getenv("DB_NAME"),
	}
	cfg.Server = models.ServerConfig{
		Port: os.Getenv("SERVER_PORT"),
	}

	return &cfg, nil
}
