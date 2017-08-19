package test

import (
	"bufio"
	"errors"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gotoeveryone/golang/common"
	"github.com/gotoeveryone/golang/common/logs"
)

// TestLog ログ出力確認
func TestLog(t *testing.T) {
	common.LoadConfig()
	key := "テスト" + time.Now().Format("20060102150405")
	logs.Info(key)

	// ファイルの存在確認
	if _, err := os.Stat(common.AppConfig.Log.Path); err != nil {
		panic(err)
	}

	// 出力先のファイルを開く
	file, err := os.Open(common.AppConfig.Log.Path)
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
