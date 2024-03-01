package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"time"
)

type Config struct {
	Host            string
	Port            uint
	DB              string
	User            string
	Password        string
	MaxConns        int32
	MinConns        int32
	MaxConnLifetime time.Duration
	MaxConnIdleTime time.Duration
}

func NewPostgresPool(cfg Config) (pool *pgxpool.Pool, err error) {
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, err
	}

	config.ConnConfig.Host = cfg.Host
	config.ConnConfig.Port = uint16(cfg.Port)
	config.ConnConfig.Database = cfg.DB
	config.ConnConfig.User = cfg.User
	config.ConnConfig.Password = cfg.Password

	// MaxConns = 25
	// MinConns = 2
	// MaxConnLifetime = 120000 * time.Millisecond
	// MaxConnIdleTime = 5 * time.Second
	config.MaxConns = cfg.MaxConns
	config.MinConns = cfg.MinConns
	config.MaxConnLifetime = cfg.MaxConnLifetime
	config.MaxConnIdleTime = cfg.MaxConnIdleTime

	pool, err = pgxpool.ConnectConfig(context.Background(), config)
	return
}
