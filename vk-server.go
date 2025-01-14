package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

const (
	vkConfirmationCode = "ВАШ_КОД_ПОДТВЕРЖДЕНИЯ" // Код подтверждения для VK Callback API
	vkToken            = "ВАШ_ТОКЕН_ДОСТУПА"     // Токен доступа VK API
	openaiAPIKey       = "ВАШ_API_КЛЮЧ_CHATGPT"   // API-ключ OpenAI
)

var (
	wg sync.WaitGroup // Для управления потоками
)

// VKMessage структура входящего сообщения от VK
type VKMessage struct {
	Type    string `json:"type"`
	GroupID int    `json:"group_id"`
	Object  struct {
		Message struct {
			FromID int    `json:"from_id"`
			Text   string `json:"text"`
		} `json:"message"`
	} `json:"object"`
}

// ChatGPTRequest структура запроса к ChatGPT API
type ChatGPTRequest struct {
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
	Model     string `json:"model"`
}

// ChatGPTResponse структура ответа от ChatGPT API
type ChatGPTResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

// sendToVK отправляет сообщение пользователю через VK API
func sendToVK(userID int, message string) {
	url := "https://api.vk.com/method/messages.send"
	params := map[string]interface{}{
		"user_id":   userID,
		"random_id": 0,
		"message":   message,
		"access_token": vkToken,
		"v":         "5.131",
	}

	body, err := json.Marshal(params)
	if err != nil {
		log.Printf("Ошибка сериализации параметров VK: %v", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Ошибка отправки сообщения VK: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("VK API вернул ошибку: %s", resp.Status)
	}
}

// processMessage обрабатывает входящее сообщение, отправляя его в ChatGPT
func processMessage(userID int, text string) {
	defer wg.Done()

	// Подготовка запроса к ChatGPT
	requestBody := ChatGPTRequest{
		Prompt:    text,
		MaxTokens: 100,
		Model:     "text-davinci-003",
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		log.Printf("Ошибка сериализации запроса к ChatGPT: %v", err)
		return
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Ошибка создания запроса к ChatGPT: %v", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+openaiAPIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Ошибка выполнения запроса к ChatGPT: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("ChatGPT API вернул ошибку: %s", resp.Status)
		return
	}

	var chatGPTResponse ChatGPTResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatGPTResponse); err != nil {
		log.Printf("Ошибка декодирования ответа ChatGPT: %v", err)
		return
	}

	// Отправка ответа обратно в VK
	if len(chatGPTResponse.Choices) > 0 {
		sendToVK(userID, chatGPTResponse.Choices[0].Text)
	} else {
		sendToVK(userID, "Извините, я не смог придумать ответ 😔")
	}
}

// handler основной обработчик входящих запросов
func handler(w http.ResponseWriter, r *http.Request) {
	var message VKMessage

	// Декодируем входящий запрос
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Printf("Ошибка декодирования запроса: %v", err)
		return
	}

	// Обрабатываем подтверждение сервера
	if message.Type == "confirmation" {
		fmt.Fprint(w, vkConfirmationCode)
		return
	}

	// Обрабатываем новое сообщение
	if message.Type == "message_new" {
		userID := message.Object.Message.FromID
		text := message.Object.Message.Text

		log.Printf("Получено сообщение от пользователя %d: %s", userID, text)

		wg.Add(1)
		go processMessage(userID, text)
	}

	// Возвращаем ответ "ok"
	fmt.Fprint(w, "ok")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	http.HandleFunc("/", handler)

	log.Printf("Сервер запущен на порту %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
