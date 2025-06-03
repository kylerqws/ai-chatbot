package adapter

import (
	"os"
	"path/filepath"

	"github.com/kylerqws/chatbot/pkg/openai/contract/service"
	"github.com/spf13/cobra"
)

type File struct {
	*service.File
	ExecStatus bool
}

type OpenAiFileAdapter struct {
	command *cobra.Command
	files   []*File
}

func NewOpenAiFileAdapter(cmd *cobra.Command) *OpenAiFileAdapter {
	return &OpenAiFileAdapter{command: cmd}
}

func (h *OpenAiFileAdapter) Files() []*File {
	return h.files
}

func (h *OpenAiFileAdapter) ExistFiles() bool {
	return len(h.files) > 0
}

func (*OpenAiFileAdapter) WrapOpenAIFile(file *service.File) *File {
	return &File{File: file}
}

func (h *OpenAiFileAdapter) WrapOpenAIFiles(files ...*service.File) []*File {
	wraps := make([]*File, len(files))
	for i := range files {
		wraps = append(wraps, h.WrapOpenAIFile(files[i]))
	}
	return wraps
}

func (h *OpenAiFileAdapter) AddFile(file *File) {
	if file != nil {
		h.files = append(h.files, file)
	}
}

func (h *OpenAiFileAdapter) AddFiles(files ...*File) {
	for i := range files {
		h.AddFile(files[i])
	}
}

func (*OpenAiFileAdapter) FileSize(path string) int64 {
	info, err := os.Stat(path)
	if err != nil || info.IsDir() {
		return 0
	}
	return info.Size()
}

func (*OpenAiFileAdapter) FileName(path string) string {
	if path == "" {
		return "unknown"
	}
	return filepath.Base(path)
}
