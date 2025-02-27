package bot

// Client - интерфейс для всех ботов (VK, Telegram и др.).
type Client interface {
	Start() error
	SendMessage(userID int, message string) error
}
