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

	// 尝试从文件中读取配置
	file, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	// 解析 JSON 数据
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
	// 将配置对象序列化为 JSON
	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	// 写入文件
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Config written to", filename)

	return nil
}

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

func getUserHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}
