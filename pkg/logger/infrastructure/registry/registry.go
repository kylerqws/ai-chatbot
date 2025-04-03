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

func (r *loggerRegistry) Register(name string, logger ctrlog.Logger) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.loggers[name] = logger
}

func (r *loggerRegistry) Logger(name string) (ctrlog.Logger, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("required logger name is not set")
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	if logger, ok := r.loggers[name]; ok {
		return logger, nil
	}

	return nil, fmt.Errorf("unregistered logger: %q", name)
}
