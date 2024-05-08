package config

import (
	"fmt"
	"path/filepath"
)

func GetOptions() ([]Option, error) {
	homeDir, err := getUserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		return []Option{}, err
	}
	configFile := filepath.Join(homeDir, ".fastGit.json")
	config, err := readConfigFromJSON(configFile)
	if err != nil {
		fmt.Println("Error reading config from JSON file:", err)
		options := GetDefaultOptions()
		config = Config{Options: options}

		err := writeConfigToJSON(configFile, config)
		if err != nil {
			fmt.Println("Error writing default config to JSON file:", err)
			return config.Options, err
		}
	}

	return config.Options, nil
}
