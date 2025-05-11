package config

import (
	"encoding/base64"
	"encoding/json"
	"os"
)

func Bootstrap[T any](configPath string) (*T, error) {
	var config T
	value := os.Getenv("CONFIG")
	if value != "" {
		decode, err := base64.StdEncoding.DecodeString(value)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(decode, &config)
	} else {
		contents, err := os.ReadFile(configPath)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(contents, &config)
	}
	return &config, nil
}
