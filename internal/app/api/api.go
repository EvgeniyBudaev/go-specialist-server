package api

import (
	"github.com/EvgeniyBudaev/go-specialist-server/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Инициализация ядра сервера

// Base API server instance description
type API struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage // Добавление поля для работы с хранилещем
}

// API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start http server/configure Loggers, router, database connection, etc
func (api *API) Start() error {
	// Trying to configure logger
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	// Подтвержение того, что logger сконфигурирован
	api.logger.Info("starting api server at port: ", api.config.BindAddr)
	// Конфигурируем маршрутизатор
	api.configureRouterField()
	// Конфигурируем хранилище
	if err := api.configureStorageField(); err != nil {
		return err
	}
	// На этапе валидного завершения стартуем http-сервер
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
