package config

import (
	"gopkg.in/yaml.v3"
	"infotecs/internal/model"
	"os"
)

func LoadConfigFromYaml() (*model.Config, error) {
	var cfg *model.Config

	file, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
