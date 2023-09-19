package storage

import (
	"database/sql"
	_ "github.com/lib/pq" // для того, чтобы отработала функция init()
	"log"
)

// Instance of storage
type Storage struct {
	config *Config
	// Database FileDescriptor
	db                *sql.DB
	userRepository    *UserRepository    // Subfield for repository interfacing (model user)
	articleRepository *ArticleRepository // Subfield for repository interfacing (model article)
}

// Storage constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

// Open connection method
func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURI) // Проверка доступности БД
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil { // Устанавливаем соединение с БД
		return err
	}
	storage.db = db
	log.Println("database connection created successfully!")
	return nil
}

// Close connection
func (storage *Storage) Close() {
	storage.db.Close()
}

// Public Repository for User
func (s *Storage) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		storage: s,
	}
	return s.userRepository
}

// Public Repository for Article
func (s *Storage) Article() *ArticleRepository {
	if s.articleRepository != nil {
		return s.articleRepository
	}
	s.articleRepository = &ArticleRepository{
		storage: s,
	}
	return s.articleRepository
}
