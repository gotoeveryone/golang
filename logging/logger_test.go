package logging

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/gotoeveryone/golib/config"
)

func TestNewLogger(t *testing.T) {
	l, err := NewLogger(config.Log{
		Prefix: "test",
	})
	if err != nil {
		t.Errorf("NewLogger error: %s", err)
	}
	if l.Level() != LevelDebug {
		t.Errorf("Level is not default level (DEBUG)")
	}

	if l.last.message != "" {
		t.Errorf("Log have do not output yet, but last message was not empty: %s", l.last.message)
	}

	m := "Test message"
	l.Info(m)

	if !strings.Contains(l.last.message, m) {
		t.Errorf("Last message not contains: %s", m)
	}

	if !strings.Contains(l.last.message, l.prefix) {
		t.Errorf("Last message not has prefix: %s", m)
	}
}

func TestNewLoggerWithLevel(t *testing.T) {
	l, err := NewLogger(config.Log{})
	if err != nil {
		t.Errorf("NewLogger error: %s", err)
	}

	lv := map[logLevel]string{
		logDebug:   "Debug",
		logInfo:    "Info",
		logWarning: "Warning",
		logError:   "Error",
	}
	for k, v := range lv {
		reflect.ValueOf(l).MethodByName(v).Call([]reflect.Value{reflect.ValueOf(fmt.Sprintf("test %s", v))})
		if l.last.level != k {
			t.Errorf("Log level not matched, actual: %d, expected: %d", l.last.level, k)
		}
		msg := string(l.fetchLevel(k))
		if !strings.Contains(l.last.message, msg) {
			t.Errorf("Log message not level contains: %s", msg)
		}
	}
}

func TestNewLoggerWithFile(t *testing.T) {
	current, _ := filepath.Abs(".")
	fn := path.Join(current, "..", "tmp", "test.log")
	if err := os.Remove(fn); err != nil {
		if !os.IsNotExist(err) {
			t.Errorf("Remove log error: %s", err)
		}
	}

	l, err := NewLogger(config.Log{
		Prefix: "TestNewLoggerWithFile",
		Path:   fn,
		Type:   "file",
	})
	if err != nil {
		t.Errorf("NewLogger error: %s", err)
	}

	m := fmt.Sprintf("test: %s", time.Now().Format("2006-01-02 15:04:05"))
	l.Info(m)

	b, err := ioutil.ReadFile(fn)
	if err != nil {
		t.Errorf("Open log error: %s", err)
	}
	if !strings.Contains(string(b), m) {
		t.Errorf("Output file message not contains: %s", m)
	}
	if !strings.Contains(string(b), l.prefix) {
		t.Errorf("Last message not has prefix: %s", l.prefix)
	}
	if !strings.Contains(string(b), l.last.message) {
		t.Errorf("Last message not has prefix: %s", l.prefix)
	}
}
