package config

import (
	"fmt"

	"github.com/KevinYouu/fastGit/functions/logs"
)

// get the default options
func GetDefaultOptions() []Option {
	return []Option{
		{Label: "fix", Value: "fix"},
		{Label: "feat", Value: "feat"},
		{Label: "refactor", Value: "refactor"},
		{Label: "chore", Value: "chore"},
		{Label: "style", Value: "style"},
		{Label: "docs", Value: "docs"},
		{Label: "build", Value: "build"},
		{Label: "revert", Value: "revert"},
		{Label: "test", Value: "test"},
	}
}

// get the options
func GetOptions() ([]Option, error) {
	config, err := getConfig()
	if err != nil {
		config.Options = GetDefaultOptions()

		if err := writeJSONConfig(config); err != nil {
			fmt.Println("Error writing default config to JSON file:", err)
			logs.Info("touching the config file")
			return nil, err
		}
	}

	return config.Options, nil
}
