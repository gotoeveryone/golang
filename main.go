package golib

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"

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
	configDir = flag.String("conf", "", "config.json at directory")
)

// LoadConfig 設定をJSONファイルから読み込む
func LoadConfig() {
	flag.Parse()
	// デフォルトは実行ファイルと同じディレクトリ
	if configDir == nil {
		(*configDir) = path.Dir(os.Args[0])
	}
	jsonValue, err := ioutil.ReadFile(*configDir + "config.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(jsonValue, &AppConfig); err != nil {
		log.Fatal(err)
	}
	// ログ初期設定
	logConfig := AppConfig.Log
	if err := logs.Init(logConfig.Prefix, logConfig.Path, logConfig.Level); err != nil {
		log.Fatal(err)
	}
}
