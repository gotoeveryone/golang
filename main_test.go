package golib

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gotoeveryone/golib/config"
)

func TestLoadConfigNoSuchFile(t *testing.T) {
	current, _ := filepath.Abs(".")

	if err := os.Remove(filepath.Join(current, "config.json")); err != nil {
		if !os.IsNotExist(err) {
			t.Errorf("Remove log error: %s", err)
		}
	}

	var config config.Config
	if err := LoadConfig(&config, ""); err != nil {
		if !strings.Contains(err.Error(), "Read file Error") {
			t.Error(err)
		}
	}
}

func TestLoadConfigParseError(t *testing.T) {
	current, _ := filepath.Abs(".")
	dir := filepath.Join(current, "hoge")

	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Error(err)
	}

	if err := ioutil.WriteFile(filepath.Join(dir, "config.json"), []byte("hoge"), 0755); err != nil {
		t.Error(err)
	}

	var config config.Config
	if err := LoadConfig(&config, dir); err != nil {
		if !strings.Contains(err.Error(), "Unmarshal error") {
			t.Error(err)
		}
	}

}

func TestLoadConfig(t *testing.T) {
	current, _ := filepath.Abs(".")

	f, err := ioutil.ReadFile(filepath.Join(current, "config.json.example"))
	if err != nil {
		t.Error(err)
	}

	if err := ioutil.WriteFile(filepath.Join(current, "config.json"), f, 0755); err != nil {
		t.Error(err)
	}

	var config config.Config
	if err := LoadConfig(&config, ""); err != nil {
		t.Error(err)
	}
}

func TestLoadConfigCustomPath(t *testing.T) {
	current, _ := filepath.Abs(".")

	f, err := ioutil.ReadFile(filepath.Join(current, "config.json.example"))
	if err != nil {
		t.Error(err)
	}

	if err := ioutil.WriteFile(filepath.Join(current, "tmp", "config.json"), f, 0755); err != nil {
		t.Error(err)
	}

	var config config.Config
	if err := LoadConfig(&config, filepath.Join(current, "tmp")); err != nil {
		t.Error(err)
	}
}
