package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

const (
	Path = "./config.yml"
)

type Config struct {
	Agent  Agent
	Server Server
}

type Agent struct {
	PollInterval   int `yaml:"poll_interval"`
	ReportInterval int `yaml:"report_interval"`
	Timeout        int `yaml:"timeout"`
}

type Server struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
	IdleTimeout  int    `yaml:"idle_timeout"`
}

func Load(path string) (Config, error) {
	var config Config
	var content []byte

	content, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
