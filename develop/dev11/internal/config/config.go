package config

import (
	"dev11/internal/config/helper"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// HttpServer is a struct that contains ip and port for http server
type HttpServer struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}

// Configuration is a struct that contains http server struct
type Configuration struct {
	HttpServer HttpServer `yaml:"http_server"`
}

// ReadConfigYaml is a function that stores config from file
func ReadConfigYaml(filePath string) (cfg Configuration, err error) {
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return cfg, err
	}

	defer helper.Closer(file)

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(&cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
