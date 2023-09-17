package api

import "github.com/EvgeniyBudaev/go-specialist-server/storage"

// Конфигурирование API сервера

// General instance for API server of REST application
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LoggerLevel string `toml:"logger_level"`
	Storage     *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
