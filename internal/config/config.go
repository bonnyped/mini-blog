package config

import (
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DbServer   DBServer   `yaml:"db_server"`
	HttpServer HTTPServer `yaml:"http_server"`
}

type DBServer struct {
	Port     string `yaml:"port" env:"PORT" env-default:"5432"`
	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
	Name     string `yaml:"name" env-default:"postgres"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env:"PASSWORD" env-default:""`
}

type HTTPServer struct {
	Port string `yaml:"port" env:"PORT" env-default:"8080"`
	Host string `yaml:"host" env:"HOST" env-default:"localhost"`
}

func MustLoad(logger *slog.Logger) *Config {
	const op = "config.MustLoad"

	var cfg Config

	err := cleanenv.ReadConfig("../config/config.yaml", &cfg)
	if err != nil {
		logger.Error("%s %w", op, err)
	}

	logger.Info("Configuration loaded", slog.Any("config", cfg))

	return &cfg
}
