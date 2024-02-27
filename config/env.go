package config

import (
	"log"
	"os"
	"strconv"
)

type Env struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPass         string `mapstructure:"DB_PASS"`
	DBName         string `mapstructure:"DB_NAME"`
	DBPort         string `mapstructure:"DB_PORT"`
}

func NewEnv() *Env {

	contextTimeoutStr := os.Getenv("CONTEXT_TIMEOUT")
	contextTimeout, err := strconv.Atoi(contextTimeoutStr)
	if err != nil {
		log.Fatalf("Erro ao converter CONTEXT_TIMEOUT para int: %v", err)
	}

	env_prod := Env{
		AppEnv:         os.Getenv("APP_ENV"),
		DBHost:         os.Getenv("DB_HOST"),
		DBUser:         os.Getenv("DB_USER"),
		DBPass:         os.Getenv("DB_PASS"),
		DBName:         os.Getenv("DB_NAME"),
		ServerAddress:  os.Getenv("SERVER_ADDRESS"),
		ContextTimeout: contextTimeout,
		DBPort:         os.Getenv("DB_PORT"),
	}
	return &env_prod

}
