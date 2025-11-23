package tests

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
)

type Config struct {
	Env  string `env:"ENV" env-default:"local"`
	Host string `env:"APP_HOST" env-default:"app-test"`
	Port int    `env:"APP_PORT" env-default:"8080"`
	DB   DB
}

type DB struct {
	User     string `env:"POSTGRES_USER" env-default:"root"`
	Password string `env:"POSTGRES_PASSWORD" env-default:"root"`
	Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
	Port     int    `env:"POSTGRES_PORT" env-default:"5432"`
	DBName   string `env:"POSTGRES_DB" env-default:"revass"`
	Schema   string `env:"POSTGRES_SCHEMA" env-default:"revass"`
}

func ReadTestConfig() Config {
	var cfg Config

	if err := cleanenv.ReadConfig("../.env", &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err.Error())
	}

	return cfg
}
