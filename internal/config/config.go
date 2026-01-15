package config

import (
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DBCOnnectionString string     `yaml:"db_connect" env-required:"true"`
	HttpServer         HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address string `yaml:"address" env-required:"true" env-default:"localhost:8080"`
}

func MustLoad(logger *slog.Logger) *Config {
	const op = "config.MustLoad"

	var cfg Config

	err := cleanenv.ReadConfig("../config/config.yaml", &cfg)
	if err != nil {
		logger.Error("%s %w", op, err)
	}

	return &cfg
}
