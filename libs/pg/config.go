package pg

import (
	"time"
)

type Config struct {
	URI             string        `env:"DB_URI"`
	Host            []string      `env:"DB_HOST" envSeparator:","`
	Database        string        `env:"DB_DATABASE"`
	Username        string        `env:"DB_USERNAME"`
	Password        string        `env:"DB_PASSWORD"`
	AppName         string        `env:"DB_APP_NAME"`
	ConnectTimeout  time.Duration `env:"DB_CONNECT_TIMEOUT"`
	MaxConns        int           `env:"DB_MAX_CONNS"`
	MinConns        int           `env:"DB_MIN_CONNS"`
	MaxConnIdleTime time.Duration `env:"DB_MAX_CONN_IDLE_TIME"`
	MaxConnLifeTime time.Duration `env:"DB_MAX_CONN_LIFE_TIME"`
	UseTLS          bool          `env:"DB_USE_TLS"`
}
