package models

import (
	"fmt"
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Environment struct {
	PostgresUser     string `env:"USERNAME_DB,required,notEmpty"`
	PostgresPassword string `env:"PASSWORD_DB,required,notEmpty"`
}

const path = "env.local"

func LoadEnv() *Environment {
	environmentVariables := Environment{}

	if err := godotenv.Load(path); err != nil {
		panic(fmt.Sprintf("Failed to load env.local file: %s", err.Error()))
	}

	if err := env.Parse(&environmentVariables); err != nil {
		panic(fmt.Sprintf("Failed to parse environment variables: %s", err.Error()))
	}

	return &environmentVariables
}
