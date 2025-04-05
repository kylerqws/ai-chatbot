package contract

import ctrsrv "github.com/kylerqws/chatbot/pkg/openai/contract/service"

type OpenAI interface {
	FileService() ctrsrv.FileService
}
