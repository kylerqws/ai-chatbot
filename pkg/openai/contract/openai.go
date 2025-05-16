package contract

import ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"

type OpenAI interface {
	FileService() ctrsvc.FileService
	JobService() ctrsvc.JobService
	ChatService() ctrsvc.ChatService
}
