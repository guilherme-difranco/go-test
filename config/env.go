package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPass         string `mapstructure:"DB_PASS"`
	DBName         string `mapstructure:"DB_NAME"`
}

func NewEnv() *Env {
	env := Env{}
	setupViper()

	err := viper.Unmarshal(&env)
	if err != nil {
		log.Fatalf("Environment can't be loaded: %v", err)
	}

	return &env
}

func setupViper() {
	viper.AutomaticEnv() // Habilita a leitura de variáveis de ambiente

	// Este bloco é opcional e só é necessário se você ainda quer suportar a leitura de um arquivo .env para desenvolvimento local
	if os.Getenv("ENV") == "development" {
		viper.SetConfigType("env")
		viper.SetConfigName(".env")
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err == nil {
			log.Println("Using .env file for configurations")
		}
	}
}
