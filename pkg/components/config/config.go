package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/KevinYouu/fastGit/pkg/components/logs"
)

type Option struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Usage int    `json:"usage"`
}

type Config struct {
	Options []Option `json:"options"`
	Patch   int8     `json:"patch"`
}

// readJSONConfig reads the config from a JSON file
func readJSONConfig(configFile string, config interface{}) error {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, config)
}

// get the config file
func getConfigFile() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	return filepath.Join(homeDir, ".fastGit.json")
}

// writeJSONConfig writes the config to a JSON file
func writeJSONConfig(config interface{}) error {
	configFile := getConfigFile()
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configFile, data, 0644)
}

// get the config
func getConfig() (Config, error) {
	configFile := getConfigFile()

	var config Config
	if err := readJSONConfig(configFile, &config); err != nil {
		fmt.Println("Error getting config from JSON file:", err)
		logs.Info("touching the config file")
		return Config{}, err
	}

	return config, nil
}
