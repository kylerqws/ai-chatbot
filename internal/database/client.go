package database

import (
	"fmt"
	"os"
)

// DatabaseClient - интерфейс для всех хранилищ данных.
type DatabaseClient interface {
	Connect() error
	Close() error
	SaveMessage(userID int, message, response string) error
}

// NewDatabaseClient создаёт нужный тип хранилища (JSON, SQLite).
func NewDatabaseClient() (DatabaseClient, error) {
	dbType := os.Getenv("DB_TYPE") // json, sqlite

	switch dbType {
	case "json":
		filePath := os.Getenv("DB_PATH")
		if filePath == "" {
			filePath = "data/messages.json"
		}
		return NewJSONDatabase(filePath)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}
