package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Contexts map[string]Context `json:"contexts"`
}

type Context struct {
	APIHost string `json:"host"`
	Auth    string `json:"auth"`
}

const configFileName = ".wsk-use"

func Read() (_ *Config, e error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("read user home directory: %v", err)
	}

	configFile, err := os.Open(filepath.Join(homeDir, configFileName))
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("open file: %v", err)
	}

	defer func() {
		if err := configFile.Close(); err != nil && e == nil {
			e = fmt.Errorf("close file: %v", err)
		}
	}()

	var config Config

	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		return nil, fmt.Errorf("decode file: %v", err)
	}

	return &config, nil
}
