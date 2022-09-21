package logutil

import (
	"fmt"
	"sync"
	"time"
)

type Logger interface {
	Log(LogLevel, string) error
}

type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelPanic
	LogLevelError
)

var LogLevelLabels = [...]string{
	LogLevelDebug:   "DEBUG",
	LogLevelInfo:    "INFO",
	LogLevelWarning: "WARNING",
	LogLevelError:   "ERROR",
	LogLevelPanic:   "PANIC",
}

type DefaultLogger struct {
	mu sync.RWMutex
}

func (l *DefaultLogger) Log(lvl LogLevel, str string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	_, err := fmt.Println(time.Now().Format(time.RFC3339) + ": " + LogLevelLabels[lvl] + ": " + str)
	return err
}
