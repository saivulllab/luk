package config

import "log"

type ConfigLoader interface {
	LoadConfig() error
	GetServiceName() string
	GetServiceVersion() string
}

func Build(configLoader ConfigLoader) {
	if err := configLoader.LoadConfig(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
}
