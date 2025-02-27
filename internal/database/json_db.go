package database

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// Message - структура для хранения сообщений.
type Message struct {
	UserID   int    `json:"user_id"`
	Message  string `json:"message"`
	Response string `json:"response"`
}

// JSONDatabase - реализация DatabaseClient для хранения в JSON-файле.
type JSONDatabase struct {
	filePath string
	mu       sync.Mutex
}

// NewJSONDatabase создаёт хранилище JSON.
func NewJSONDatabase(filePath string) (*JSONDatabase, error) {
	db := &JSONDatabase{filePath: filePath}

	// Создаём файл, если его нет
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := db.save([]Message{}) // Создаём пустой JSON-файл
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

// Connect - заглушка (для совместимости с интерфейсом).
func (j *JSONDatabase) Connect() error {
	return nil
}

// Close - заглушка (файл не требует закрытия).
func (j *JSONDatabase) Close() error {
	return nil
}

// SaveMessage - сохраняет сообщение в JSON-файл.
func (j *JSONDatabase) SaveMessage(userID int, message, response string) error {
	j.mu.Lock()
	defer j.mu.Unlock()

	// Загружаем существующие данные
	messages, err := j.load()
	if err != nil {
		return fmt.Errorf("failed to load JSON data: %v", err)
	}

	// Добавляем новое сообщение
	messages = append(messages, Message{
		UserID:   userID,
		Message:  message,
		Response: response,
	})

	// Сохраняем обратно в файл
	return j.save(messages)
}

// load загружает данные из JSON-файла.
func (j *JSONDatabase) load() ([]Message, error) {
	file, err := os.Open(j.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var messages []Message
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&messages); err != nil {
		return nil, err
	}

	return messages, nil
}

// save записывает данные в JSON-файл.
func (j *JSONDatabase) save(messages []Message) error {
	file, err := os.Create(j.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Красивый формат JSON
	return encoder.Encode(messages)
}
