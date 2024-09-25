package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

//Instance of storage

type Storage struct {
	config *Config
	//DB file descriptor
	db *sql.DB
}

// Storage constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

// Open connection method
func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	log.Println("db connection successfully!")
	return nil
}

// Close connection
func (storage *Storage) Close() {
	storage.db.Close()

}
