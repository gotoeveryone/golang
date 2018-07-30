package golib

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	configDir = flag.String("conf", "./", "config.json at directory")
)

// LoadConfig 設定をJSONファイルから構造体へ読み込む
func LoadConfig(config interface{}, customPath string) error {
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

	return nil
}
