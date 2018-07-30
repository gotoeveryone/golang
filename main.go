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
	executable, pathError = os.Executable()
	configDir             = flag.String("conf", executable, "config.json at directory")
)

// LoadConfig is read the configuration from the JSON file into the structure
func LoadConfig(config interface{}, customPath string) error {
	if pathError != nil {
		return pathError
	}

	var configPath string
	if customPath != "" {
		configPath = customPath
	} else {
		flag.Parse()
		// The default is the same directory as the executable file
		if configDir == nil {
			configPath = executable
		} else {
			configPath = (*configDir)
		}
	}

	// Read from config.json
	jsonValue, err := ioutil.ReadFile(filepath.Join(configPath, "config.json"))
	if err != nil {
		return fmt.Errorf("Read file Error: %s", err)
	}

	// Parse to JSON
	if err := json.Unmarshal(jsonValue, config); err != nil {
		return fmt.Errorf("Unmarshal error: %s", err)
	}

	return nil
}
