package golib

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/gotoeveryone/golib/config"
)

const (
	logDebug   = 1
	logInfo    = 2
	logWarning = 3
	logError   = 4

	// LevelDebug is debug level text
	LevelDebug = "DEBUG"
	// LevelInfo is info level text
	LevelInfo = "INFO"
	// LevelWarning is warning level text
	LevelWarning = "WARNING"
	// LevelError is error level text
	LevelError = "ERROR"
)

var (
	logLevels = map[config.LogLevel]logLevel{
		LevelDebug:   logDebug,
		LevelInfo:    logInfo,
		LevelWarning: logWarning,
		LevelError:   logError,
	}
)

type (
	// Logger is logger with configuration.
	Logger struct {
		prefix string
		level  logLevel
		last   last
	}
	logLevel int
	last     struct {
		level   logLevel
		message string
	}
)

// NewLogger is logger with prefix.
func NewLogger(c config.Log) (*Logger, error) {
	if c.Type == "file" {
		// Create directory when not exists specified path.
		if err := os.MkdirAll(path.Dir(c.Path), 0755); err != nil {
			return nil, errors.New("Directory can't created")
		}

		// Open log file
		f, err := os.OpenFile(c.Path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
		if err != nil {
			return nil, err
		}

		log.SetOutput(io.MultiWriter(f, os.Stdout))
	}

	level, ok := logLevels[c.Level]
	if !ok {
		level = logDebug
	}

	return &Logger{prefix: c.Prefix, level: level}, nil
}

// Debug is debug level logging
func (l *Logger) Debug(v interface{}) {
	l.out(logDebug, false, v)
}

// Info is info level logging
func (l *Logger) Info(v interface{}) {
	l.out(logInfo, false, v)
}

// Warning is info level logging
func (l *Logger) Warning(v interface{}) {
	l.out(logWarning, false, v)
}

// Error is error level logging
func (l *Logger) Error(v interface{}) {
	l.out(logError, true, v)
}

// Level is return log level
func (l *Logger) Level() config.LogLevel {
	return l.fetchLevel(l.level)
}

func (l *Logger) fetchLevel(lv logLevel) config.LogLevel {
	for k, v := range logLevels {
		if lv == v {
			return k
		}
	}
	return LevelDebug
}

func (l *Logger) isOut(target logLevel) bool {
	return l.level <= target
}

func (l *Logger) out(target logLevel, outError bool, v interface{}) {
	// Output message when specified level over
	if !l.isOut(target) {
		return
	}

	lt := l.fetchLevel(target)

	if l.prefix != "" {
		l.last = last{
			level:   target,
			message: fmt.Sprintf("[%s] %s %s", l.prefix, lt, v),
		}
	} else {
		l.last = last{
			level:   target,
			message: fmt.Sprintf("%s %s", lt, v),
		}
	}

	if outError {
		log.Println(fmt.Errorf(l.last.message))
	} else {
		log.Println(l.last.message)
	}
}
