package pg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"strings"
)

func MakePool(cfg *Config) (*pgxpool.Pool, error) {
	connCfg, err := MakePoolCfg(cfg)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	pool, err := pgxpool.NewWithConfig(ctx, connCfg)
	if err != nil {
		return nil, fmt.Errorf("error creating connection pool: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return pool, nil
}

func MakePoolCfg(cfg *Config) (*pgxpool.Config, error) {
	var dsn string

	if len(cfg.URI) > 0 {
		dsn = cfg.URI
	} else {
		dsn = fmt.Sprintf("postgres://%s:%s@%s/%s?connect_timeout=%d&application_name=%s&sslmode=%s",
			cfg.Username, cfg.Password, strings.Join(cfg.Host, ","),
			cfg.Database, cfg.ConnectTimeout.Seconds(),
			cfg.AppName, tlsMode(cfg.UseTLS))
	}

	connCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("error parsing pool config: %w", err)
	}

	connCfg.MaxConns = int32(cfg.MaxConns)
	connCfg.MinConns = int32(cfg.MinConns)
	connCfg.MaxConnIdleTime = cfg.MaxConnIdleTime
	connCfg.MaxConnLifetime = cfg.MaxConnLifeTime

	return connCfg, nil
}

// tlsMode returns the corresponding value for the sslMode parameter
func tlsMode(useTLS bool) string {
	if useTLS {
		return "require"
	}

	return "disable"
}
