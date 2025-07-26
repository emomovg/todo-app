package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

type PostgresConfig struct {
	Host          string
	Port          int
	Username      string
	Password      string
	DBName        string
	SSLMode       string
	DBMaxCons     int
	DBMinCons     int
	DBMaxLifetime string
	DBMaxIdleTime string
}

func New(ctx context.Context, cfg PostgresConfig) (*Postgres, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)

	poolCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Pool config: %w", err)
	}

	poolCfg.MaxConns = int32(cfg.DBMaxCons)
	poolCfg.MinConns = int32(cfg.DBMinCons)

	if maxLifetime, err := time.ParseDuration(cfg.DBMaxLifetime); err == nil {
		poolCfg.MaxConnLifetime = maxLifetime
	}

	if maxIdleTime, err := time.ParseDuration(cfg.DBMaxIdleTime); err == nil {
		poolCfg.MaxConnIdleTime = maxIdleTime
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create Pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("db ping failed: %w", err)
	}

	return &Postgres{Pool: pool}, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
