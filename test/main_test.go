package test

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gotoeveryone/golib"
	"github.com/gotoeveryone/golib/logs"
)

// TestLog ログ出力確認
func TestLog(t *testing.T) {
	config := golib.Config{}
	if err := golib.LoadConfig(&config, ""); err != nil {
		panic(errors.New("LoadConfig error"))
	}

	// ログ書き込み
	key := "テスト" + time.Now().Format("20060102150405")
	logs.Info(key)

	// ファイルの存在確認
	if _, err := os.Stat(config.Log.Path); err != nil {
		panic(err)
	}

	// 出力先のファイルを開く
	file, err := os.Open(config.Log.Path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 出力した文字列が含まれているか
	s := bufio.NewScanner(file)
	result := false
	for s.Scan() {
		if strings.Contains(s.Text(), key) {
			result = true
		}
	}
	if !result {
		panic(errors.New("Not Contains"))
	}
}

// TestAppConfig 内部保持している設定ファイルで読み出す
func TestAppConfig(t *testing.T) {
	if err := golib.LoadConfig(nil, ""); err != nil {
		panic(fmt.Errorf("LoadConfig error: %s", err))
	}
	if &golib.AppConfig == nil {
		panic(fmt.Errorf("AppConfig not initialized"))
	}
}
