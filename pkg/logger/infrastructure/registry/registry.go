package registry

import (
	"fmt"
	"strings"
	"sync"

	ctrlog "github.com/kylerqws/chatbot/pkg/logger/contract/logger"
	ctrreg "github.com/kylerqws/chatbot/pkg/logger/contract/registry"
)

type loggerRegistry struct {
	mu      sync.RWMutex
	loggers map[string]ctrlog.Logger
}

func New() ctrreg.LoggerRegistry {
	return &loggerRegistry{
		loggers: make(map[string]ctrlog.Logger),
	}
}

func (r *loggerRegistry) Register(name string, log ctrlog.Logger) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.loggers[name] = log
}

func (r *loggerRegistry) Logger(name string) (ctrlog.Logger, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("required logger name is not set")
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	if log, ok := r.loggers[name]; ok {
		return log, nil
	}

	return nil, fmt.Errorf("unregistered logger: %q", name)
}
