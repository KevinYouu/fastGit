package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/user"
)

type Option struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
type Config struct {
	Options []Option `json:"options"`
}

func readConfigFromJSON(filename string) (Config, error) {
	var config Config

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	// Read the file
	data, err := io.ReadAll(file)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func writeConfigToJSON(filename string, config Config) error {
	// Marshal the config
	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	// Write the file
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Config written to", filename)

	return nil
}

// getUserHomeDir gets the user's home directory
func getUserHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}
