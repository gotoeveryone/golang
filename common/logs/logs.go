package logs

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

var (
	prefix string
	level  int
	levels = map[string]int{
		"DEBUG": 0,
		"INFO":  1,
		"ERROR": 2,
		"FATAL": 3,
	}
	initialize = false
)

// Init ログ初期設定を行う
func Init(pre, p, l string) error {
	if initialize {
		return errors.New("Already initialized")
	}
	// ディレクトリがなければ生成
	if err := os.MkdirAll(path.Dir(p), os.ModeDir); err != nil {
		return errors.New("Directory can't created")
	}
	// ログ出力設定
	file, err := os.OpenFile(p, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}

	log.SetOutput(io.MultiWriter(file, os.Stdout))
	prefix = pre
	level = levels[l]
	initialize = true
	return nil
}

// Debug デバッグ
func Debug(v ...interface{}) {
	logging("DEBUG", false, v)
}

// Info 通知
func Info(v ...interface{}) {
	logging("INFO", false, v)
}

// Error エラー
func Error(v ...interface{}) {
	logging("ERROR", true, v)
}

// Fatal 致命的
func Fatal(v ...interface{}) {
	logging("FATAL", true, v)
}

func logging(target string, outError bool, v interface{}) {
	if !initialize {
		log.Fatal("Log not initialized")
	}
	// 指定レベル以上の場合のみ出力
	if level <= levels[target] {
		if prefix == "" {
			target = prefix + " " + target
		}

		// エラー出力するかどうか
		var out interface{}
		if outError {
			out = fmt.Errorf("%s %s", target, v)
		} else {
			out = fmt.Sprintf("%s %s", target, v)
		}

		// FATALの場合はログ出力して終了
		if target == "FATAL" {
			log.Fatalln(out)
		} else {
			log.Println(out)
		}
	}
}
