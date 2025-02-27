# ChatGPT Bot Server

ChatGPT Bot Server – это сервер чат-бота, который интегрируется с различными платформами (VK, Telegram) и использует разные AI-провайдеры (например, OpenAI).

## Возможности
- Поддержка VK и Telegram
- Выбор AI-провайдера (OpenAI, можно добавить другие)
- Гибкое хранилище данных (JSON-файл или SQLite)
- Запуск через CLI

## 1. Установка

### Клонируем репозиторий
git clone https://github.com/kylerqws/chatgpt-bot.git
cd chatgpt-bot

### Устанавливаем зависимости
go mod tidy

## 2. Настройка

### Создаём .env файл
Создай .env в корневой папке и укажи настройки:

AI_PROVIDER=openai
OPENAI_API_KEY=your_openai_key

DB_TYPE=json
DB_PATH=data/messages.json

DB_TYPE=sqlite
DB_PATH=data/chatgpt-bot.db

VK_TOKEN=your_vk_token
TELEGRAM_TOKEN=your_telegram_token

SERVER_PORT=5000

## 3. Запуск сервера

### Запускаем бота
go run main.go serve

Теперь бот работает с VK и Telegram (если заданы токены).

## 4. Управление обучением AI

### Создать задачу на обучение
go run main.go train create --file-id=file-123456 --model=gpt-3.5-turbo

### Проверить статус обучения
go run main.go train status job-123456

### Отменить обучение
go run main.go train cancel job-123456

## 5. Проверка сохранённых сообщений

Если используется JSON-хранилище, можно проверить сохранённые данные:
cat data/messages.json

Если используется SQLite, можно открыть БД:
sqlite3 data/chatgpt-bot.db

И выполнить SQL-запрос:
SELECT * FROM messages;

## 6. Добавление новых AI-сервисов

Этот проект поддерживает разных AI-провайдеров.
Чтобы добавить новый сервис (например, Azure OpenAI):

1. Создай новый файл internal/ai/azure.go.
2. Реализуй интерфейс AIClient.
3. Добавь его в NewAIClient() в client.go.
4. Теперь можно выбрать AI_PROVIDER=azure в .env!

## 7. Лицензия

Этот проект распространяется под MIT License.
Используйте и улучшайте!

## Итог

- README.md описывает установку, настройку и запуск
- Понятно, как использовать VK, Telegram и OpenAI
- Есть инструкции по обучению AI
- Можно легко менять хранилище и AI-сервис

Теперь всё задокументировано!
