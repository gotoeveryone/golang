package logs

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	level  int
	levels = map[string]int{
		"INFO":  1,
		"ERROR": 2,
		"FATAL": 3,
	}
	initialize = false
)

// InitLog ログ初期設定を行う
func Init(p, l string) error {
	if initialize {
		return errors.New("ログ設定はすでに初期化済みです。")
	}
	// ログ出力設定
	file, err := os.OpenFile(p, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	log.SetOutput(io.MultiWriter(file, os.Stdout))
	level = levels[l]
	initialize = true
	return nil
}

// Info 通知
func Info(v interface{}) {
	if level >= levels["INFO"] {
		log.Println(fmt.Sprintf("%s %s", "INFO", v))
	}
}

// Error エラー
func Error(v interface{}) {
	if level >= levels["ERROR"] {
		log.Println(fmt.Errorf("%s %s", "ERROR", v))
	}
}

// Fatal 致命的
func Fatal(v interface{}) {
	if level >= levels["FATAL"] {
		log.Fatal(fmt.Errorf("%s %s", "FATAL", v))
	}
}
