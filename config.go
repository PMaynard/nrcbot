package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	GithubToken string `json:"github_token"`
}

func ConfigFile(file string) (Config, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}
