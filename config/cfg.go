package config

import (
	"encoding/json"
	"os"
)

const path = "config/config.json"

type Config struct {
	Host    string `json:"host,omitempty"`
	Port    string `json:"port,omitempty"`
	Dbname  string `json:"dbname,omitempty"`
	SSlmode string `json:"ssl_mode,omitempty"`
}

func ReadCfg() *Config {
	file, _ := os.Open(path)
	defer closeFile(file)

	decoder := json.NewDecoder(file)

	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {

	}

	return &config
}

func closeFile(file *os.File) {
	_ = file.Close()
}
