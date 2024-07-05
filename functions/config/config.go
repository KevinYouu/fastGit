package config

import (
	"encoding/json"
	"os"
)

type Option struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
type Config struct {
	Options []Option `json:"options"`
}

// readJSONConfig reads the config from a JSON file
func readJSONConfig(configFile string, config interface{}) error {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, config)
}

// writeJSONConfig writes the config to a JSON file
func writeJSONConfig(configFile string, config interface{}) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configFile, data, 0644)
}

// getUserHomeDir gets the user's home directory
func getUserHomeDir() (string, error) {
	return os.UserHomeDir()
}
