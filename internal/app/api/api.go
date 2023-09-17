package api

// Инициализация ядра сервера

// Base API server instance description
type API struct {
	config *Config
}

// API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
	}
}

// Start http server/configure Loggers, router, database connection, etc
func (api *API) Start() error {
	return nil
}
