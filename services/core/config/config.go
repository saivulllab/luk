package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"os"
	"pg"
)

const (
	serviceName    = "LUK"
	serviceVersion = "0.1.0"
	configPath     = "/etc/luk/.env"
)

var jwtSecret []byte

type Config struct {
	Service
	Kafka
	DB pg.Config
}

type Service struct {
	Host string `env:"SERVICE_HOST"`
	Port string `env:"SERVICE_PORT"`

	jwtS []byte `env:"JWT_SECRET"`

	Env string `env:"SERVICE_ENV"`
}

type Kafka struct {
	Brokers []string `env:"KAFKA_BROKERS"`
	Topic   string   `env:"KAFKA_TOPIC"`
}

func (c *Config) LoadConfig() error {
	cfgPath := os.Getenv("LUK_CONFIG_PATH")

	if len(cfgPath) == 0 {
		cfgPath = configPath
	}

	if err := godotenv.Load(cfgPath); err != nil {
		return err
	}

	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return err
	}

	*c = cfg

	jwtSecret = c.jwtS

	return nil
}

func (c *Config) GetServiceName() string {
	return serviceName
}

func (c *Config) GetServiceVersion() string {
	return serviceVersion
}

func GetJWTSecret() []byte {
	return jwtSecret
}
