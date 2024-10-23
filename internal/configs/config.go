package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Resource struct {
	Name            string `yaml:"name"`
	Endpoint        string `yaml:"endpoint"`
	Destination_url string `yaml:"destination_url"`
}

type Configuration struct {
	Server struct {
		Host        string `yaml:"host"`
		Listen_port string `yaml:"listen_port"`
	} `yaml:"server"`
	Resources []Resource `yaml:"resources"`
}

func Load() (Configuration, error) {
	config := Configuration{}
	rootDir, err := os.Getwd()
	if err != nil {
		return config, fmt.Errorf("error getting root directory: %s", err)
	}
	configFile := filepath.Join(rootDir, "data", "config.yaml")

	data, err := os.ReadFile(configFile)
	if err != nil {
		return Configuration{}, fmt.Errorf("error loading config file: %s", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	return config, nil
}
