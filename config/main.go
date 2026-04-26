package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Carlltz/aj/utils"
)

var ConfigPath = filepath.Join(os.Getenv("HOME"), ".config", "aj", "config.json")

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = loadConfig()
	}
	return config
}

func SetConfig(cfg *Config) error {
	config = cfg
	jsonData, err := json.Marshal(*config)
	if err != nil {
		return err
	}

	err = os.WriteFile(ConfigPath, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func loadConfig() *Config {
	data, err := os.ReadFile(ConfigPath)
	if err != nil {
		if os.IsNotExist(err) {
			jsonData, err := json.Marshal(Config{
				Os: utils.GetOS(),
			})
			if err != nil {
				fmt.Println("Failed to marshall config", err)
			}
			if err := os.MkdirAll(filepath.Dir(ConfigPath), 0o755); err != nil {
				fmt.Println("Failed to create config directory", err)
			}
			err = os.WriteFile(ConfigPath, jsonData, 0644)
			if err != nil {
				fmt.Println("Failed to write config.json", err)
			}
			data = jsonData
		} else {
			fmt.Println("Error loading config.json", err)
			return &Config{}
		}
	}
	// parse data into Config struct
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		fmt.Println("Error parsing config.json", err)
		return &Config{}
	}
	return &cfg
}
