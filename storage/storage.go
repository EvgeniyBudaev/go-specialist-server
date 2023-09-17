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
	db *sql.DB
}

// Storage Constructor
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
