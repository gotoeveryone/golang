package golib

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gotoeveryone/golib/logs"
)

type (
	// Config 設定
	Config struct {
		Log   Log   `json:"log"`
		Cache Cache `json:"cache"`
		DB    DB    `json:"db"`
		Mail  Mail  `json:"mail"`
	}

	// Log ログ設定
	Log struct {
		Prefix string `json:"prefix"`
		Path   string `json:"path"`
		Level  string `json:"level"`
	}

	// Cache キャッシュサーバ接続設定
	Cache struct {
		Use  bool   `json:"use"`
		Host string `json:"host"`
		Port int    `json:"port"`
		Auth string `json:"auth"`
	}

	// DB データベース接続設定
	DB struct {
		Name     string `json:"name"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
	}

	// Mail メール接続設定
	Mail struct {
		SMTP      string   `json:"smtp"`
		Port      int      `json:"port"`
		UseTLS    bool     `json:"useTLS"`
		User      string   `json:"user"`
		Password  string   `json:"password"`
		From      string   `json:"from"`
		FromAlias string   `json:"fromAlias"`
		To        []string `json:"to"`
	}
)

var (
	// AppConfig アプリケーションが保持する設定
	AppConfig Config
	configDir = flag.String("conf", "./", "config.json at directory")
)

// LoadConfig 設定をJSONファイルから読み込む
func LoadConfig(config *Config, customPath string) error {
	// 引数の構造体がnilならデフォルトを利用
	if config == nil {
		config = &AppConfig
	}
	var configPath string
	if customPath != "" {
		configPath = customPath
	} else {
		flag.Parse()
		// デフォルトは実行ファイルと同じディレクトリ
		if configDir == nil {
			executable, _ := os.Executable()
			configPath = filepath.Dir(executable) + "/"
		} else {
			configPath = (*configDir)
		}
	}

	// 構造体読み込み
	jsonValue, err := ioutil.ReadFile(fmt.Sprintf("%sconfig.json", configPath))
	if err != nil {
		return fmt.Errorf("Read file Error: %s", err)
	}

	// JSON変換
	if err := json.Unmarshal(jsonValue, config); err != nil {
		return fmt.Errorf("Unmarshal error: %s", err)
	}

	// ログ初期設定
	logConfig := config.Log
	if err := logs.Init(logConfig.Prefix, logConfig.Path, logConfig.Level); err != nil {
		return fmt.Errorf("LogConfig error: %s", err)
	}

	return nil
}
