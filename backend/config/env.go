package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost         string `mapstructure:"DB_HOST"`
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
	viper.AutomaticEnv() // Prioriza variáveis de ambiente sobre o arquivo .env

	// Define o arquivo padrão .env para fallback
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..") // Adiciona caminhos extras caso seja necessário

	// Tenta carregar o arquivo .env, mas não falha se não for encontrado
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using .env file for configurations")
	}
}
