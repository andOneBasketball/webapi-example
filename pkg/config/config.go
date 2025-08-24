package config

import (
	"os"
	"webapi-example/pkg/models"

	"gopkg.in/yaml.v2"
)

var Cfg *models.Config

func LoadConfig(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var config models.Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return err
	}
	Cfg = &config
	return nil
}
