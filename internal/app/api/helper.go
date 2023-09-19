package api

import (
	"github.com/EvgeniyBudaev/go-specialist-server/internal/app/middleware"
	"github.com/EvgeniyBudaev/go-specialist-server/storage"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	prefix string = "/api/v1/"
)

// Пытаемся отконфигурировать наш API инстанс (а конкретнее поле logger)
func (a *API) configureLoggerField() error {
	logLevel, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(logLevel)
	return nil
}

// Пытаемся сконфигурировать маршрутизатор (а конкретнее поле router API)
func (a *API) configureRouterField() {
	a.router.HandleFunc(prefix+"/articles", a.GetAllArticles).Methods("GET")
	// Было до JWT
	// a.router.HandleFunc(prefix+"/articles/{id}", a.GetArticleById).Methods("GET")
	// Теперь требует наличие JWT
	a.router.Handle(prefix+"/articles/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.GetArticleById),
	)).Methods("GET")
	//
	a.router.HandleFunc(prefix+"/articles/{id}", a.DeleteArticleById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/articles", a.CreateArticle).Methods("POST")
	a.router.HandleFunc(prefix+"/user/register", a.RegisterUser).Methods("POST")
	// new pair for auth
	a.router.HandleFunc(prefix+"/user/auth", a.PostToAuth).Methods("POST")
}

// Пытаемся сконфигурировать хранилище (storage API)
func (a *API) configureStorageField() error {
	storageNew := storage.New(a.config.Storage)
	if err := storageNew.Open(); err != nil {
		return err
	}
	a.storage = storageNew
	return nil
}
