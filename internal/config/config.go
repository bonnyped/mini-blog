package config

import (
	"log"
	"mini-blog/internal/logger/sl"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DBCOnnectionString string     `yaml:"db_connect" env-required:"true"`
	HttpServer         HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address string `yaml:"address" env-required:"true" env-default:"localhost:8080"`
}

func MustLoad() *Config {
	const op = "config.MustLoad"

	var cfg Config

	err := cleanenv.ReadConfig("../config/config.yaml", &cfg)
	if err != nil {
		log.Fatal(sl.Err(op, err))
	}

	return &cfg
}
