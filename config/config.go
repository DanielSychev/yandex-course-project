package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"gitlab.crja72.ru/golang/2025/spring/course/students/308006-dsycev62-course-1343/postgress"
)

type Config struct {
	Postgres postgress.Config `yaml:"POSTGRES" env:"POSTGRES"`
	GRPCPort int              `yaml:"GRPC_PORT" env:"GRPC_PORT" envDefault:"5252"`
	RestPort int              `yaml:"REST_PORT" env:"REST_PORT" envDefault:"5151"`
}

func New() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig("./config/.env", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
