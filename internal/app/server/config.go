package server

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host     string `yaml:"host"`
	LogLevel string `yaml:"logLevel"`
}

func NewConfig() *Config {
	return &Config{
		Host:     ":3000",
		LogLevel: "debug",
	}
}

func (config *Config) ParseConfig(configPath string) error {

	file, err := os.Open(configPath)

	if err != nil {
		return err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return err
	}
	return nil
}
