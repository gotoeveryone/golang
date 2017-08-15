package common

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"github.com/gotoeveryone/golang/logs"
)

type (
	// Config 設定
	Config struct {
		Log   Log   `json:"log"`
		Redis Redis `json:"redis"`
		DB    DB    `json:"db"`
		Mail  Mail  `json:"mail"`
	}

	// Log ログ設定
	Log struct {
		Path  string `json:"path"`
		Level string `json:"level"`
	}

	// Redis Redis接続設定
	Redis struct {
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
	configDir = flag.String("conf", "", "config.json at directory")
)

// LoadConfig 設定をJSONファイルから読み込む
func LoadConfig(config *Config) {
	flag.Parse()
	jsonValue, err := ioutil.ReadFile(*configDir + "config.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(jsonValue, &config); err != nil {
		log.Fatal(err)
	}
	// ログ初期設定
	if err := logs.Init(config.Log.Path, config.Log.Level); err != nil {
		log.Fatal(err)
	}
}
