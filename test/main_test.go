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
	var config golib.Config
	if err := golib.LoadConfig(&config, ""); err != nil {
		panic(fmt.Errorf("LoadConfig error: %s", err))
	}

	// ログ初期設定
	timeKey := time.Now().Format("20060102150405")
	logConfig := config.Log
	logFile := logConfig.Path + timeKey + ".log"
	if err := logs.Init(logConfig.Prefix, logFile, logConfig.Level); err != nil {
		panic(fmt.Errorf("LogConfig error: %s", err))
	}

	// ログ書き込み
	key := "テスト" + time.Now().Format("20060102150405")
	logs.Info(key)

	// ファイルの存在確認
	if _, err := os.Stat(logFile); err != nil {
		panic(err)
	}

	// 出力先のファイルを開く
	file, err := os.Open(logFile)
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

// TestNotRead 設定ファイルが読み込めない
func TestNotRead(t *testing.T) {
	var config golib.Config
	// ポインタを渡さないとJSON変換に失敗する
	if err := golib.LoadConfig(config, ""); err != nil {
		if !strings.HasPrefix(err.Error(), "Unmarshal error") {
			panic(fmt.Errorf("LoadConfig error: %s", err))
		}
	}
}
