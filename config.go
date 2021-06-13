package gobootstrap

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	DbHost     string `env:"DB_HOST,required"`
	DbPort     string `env:"DB_PORT,required"`
	DbUser     string `env:"DB_USER,required"`
	DbPassword string `env:"DB_PASSWORD,required"`
	DbDatabase string `env:"DB_DATABASE,required"`
}

func ParseEnv() (Config, error) {
	_ = godotenv.Load(".env")

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
