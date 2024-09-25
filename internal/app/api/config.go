package api

import "awesomeProject/storage"

// general instance for api server of REST app
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LoggerLevel string `toml:"logger_level"`
	Storage     *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8081",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
