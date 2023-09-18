package api

import (
	"encoding/json"
	"fmt"
	"github.com/EvgeniyBudaev/go-specialist-server/internal/app/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Вспомогательная структура для формирования сообщений
type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

// Full API Handler initialization file
func initHandlers(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

// Возвращает все статьи из БД на данный момент
func (api *API) GetAllArticles(writer http.ResponseWriter, req *http.Request) {
	// Инициализируем хедеры
	initHandlers(writer)
	// Логируем момент начала обработки запроса
	api.logger.Info("get all articles GET /api/v1/articles")
	// Пытаемся что-то получить из БД
	articles, err := api.storage.Article().SelectAll()
	if err != nil {
		api.logger.Info("error while Articles.SelectAll:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "we have some troubles to accessing database. Try again later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(articles)
}

func (api *API) GetArticleById(writer http.ResponseWriter, req *http.Request) {
	initHandlers(writer)
	api.logger.Info("get article by id GET /api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("troubles while parsing {id} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "unappropriated id value. don't use ID as uncasting to int value",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	article, ok, err := api.storage.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("troubles while accessing database table (articles) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "we have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("can't find article with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "article with that ID does not exists in database",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(article)
}

func (api *API) DeleteArticleById(writer http.ResponseWriter, req *http.Request) {
	initHandlers(writer)
	api.logger.Info("delete article by id DELETE /api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("troubles while parsing {id} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "unappropriated id value. don't use ID as uncasting to int value",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok, err := api.storage.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("troubles while accessing database table (articles) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "we have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("can't find article with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "article with that ID does not exists in database",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, err = api.storage.Article().DeleteById(id)
	if err != nil {
		api.logger.Info("troubles while deleting database element from table (articles) with id. err:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "we have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(202)
	msg := Message{
		StatusCode: 202,
		Message:    fmt.Sprintf("article with ID %d successfully deleted", id),
		IsError:    false,
	}
	json.NewEncoder(writer).Encode(msg)
}

func (api *API) CreateArticle(writer http.ResponseWriter, req *http.Request) {
	initHandlers(writer)
	api.logger.Info("create article POST /api/v1/articles")
	var article models.Article
	err := json.NewDecoder(req.Body).Decode(&article)
	if err != nil {
		api.logger.Info("invalid json received from client")
		msg := Message{
			StatusCode: 400,
			Message:    "provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	a, err := api.storage.Article().Create(&article)
	if err != nil {
		api.logger.Info("troubles while creating new article:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "we have some troubles to accessing database, Try again.",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(a)
}

func (api *API) RegisterUser(writer http.ResponseWriter, req *http.Request) {
	initHandlers(writer)
	api.logger.Info("register user POST /api/v1/user/register")
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		api.logger.Info("invalid json received from client")
		msg := Message{
			StatusCode: 400,
			Message:    "provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok, err := api.storage.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("troubles while accessing database table (users) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "we have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if ok {
		api.logger.Info("user with that ID already exists")
		msg := Message{
			StatusCode: 400,
			Message:    "user with that login already exists in database",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	userAdded, err := api.storage.User().Create(&user)
	if err != nil {
		api.logger.Info("troubles while accessing database table (users) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "we have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	msg := Message{
		StatusCode: 201,
		Message:    fmt.Sprintf("user {login:%s} successfully registered!", userAdded.Login),
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}
