package api

// Конфигурирование API сервера

// General instance for API server of REST application
type Config struct {
	BindAddr string `toml:bind_addr`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
