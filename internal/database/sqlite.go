package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// SQLiteClient - реализация DatabaseClient для SQLite.
type SQLiteClient struct {
	db *sql.DB
}

// NewSQLiteClient создаёт подключение к SQLite.
func NewSQLiteClient(dbPath string) (*SQLiteClient, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open SQLite database: %v", err)
	}

	client := &SQLiteClient{db: db}
	if err := client.createTables(); err != nil {
		return nil, err
	}

	return client, nil
}

// Connect - подключение (уже выполнено в NewSQLiteClient, поэтому просто возвращаем nil).
func (s *SQLiteClient) Connect() error {
	return nil
}

// Close - закрывает соединение с БД.
func (s *SQLiteClient) Close() error {
	return s.db.Close()
}

// SaveMessage - сохраняет сообщение и ответ в БД.
func (s *SQLiteClient) SaveMessage(userID int, message, response string) error {
	_, err := s.db.Exec("INSERT INTO messages (user_id, message, response) VALUES (?, ?, ?)", userID, message, response)
	if err != nil {
		return fmt.Errorf("failed to save message: %v", err)
	}
	return nil
}

// createTables создаёт таблицы, если их нет.
func (s *SQLiteClient) createTables() error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		message TEXT NOT NULL,
		response TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := s.db.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("failed to create tables: %v", err)
	}
	return nil
}
