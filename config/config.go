package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
)

type Config struct {
	Port  int `yaml:"port" default:"8080"`
	Oauth struct {
		RedirectURL  string `yaml:"redirectUrl"`
		ClientID     string `yaml:"clientId"`
		ClientSecret string `yaml:"clientSecret"`
	} `yaml:"oauth"`
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
		Host     string `yaml:"host"`
	} `yaml:"db"`
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
	Logging struct {
		Router bool `yaml:"router"`
	}
}

func LoadConfig(log *slog.Logger) Config {
	configFileName := flag.String("config", "config.yml", "Specify config file")
	flag.Parse()
	file, err := os.ReadFile(*configFileName)
	if err != nil {
		log.Error("could not load config", configFileName)
		os.Exit(2)
	}
	config := Config{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Error("could not read config", err)
		os.Exit(2)
	}
	log.Info("configuration loaded")
	return config
}
