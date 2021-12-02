package openwhisk

import (
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Auth    string
	APIHost string
}

const configFileName = ".wskprops"

func WriteConfig(config *Config) (e error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("read user home directory: %v", err)
	}

	configFile, err := os.Create(filepath.Join(homeDir, configFileName))
	if err != nil {
		return fmt.Errorf("create configuration: %v", err)
	}

	defer func() {
		if err := configFile.Close(); err != nil && e == nil {
			e = fmt.Errorf("close configuration: %v", err)
		}
	}()

	fmt.Fprintf(configFile, "AUTH=%s\n", config.Auth)
	fmt.Fprintf(configFile, "APIHOST=%s\n", config.APIHost)

	return nil
}
