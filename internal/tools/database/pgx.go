package database

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/leguminosa/kounat/internal/tools"
	"github.com/leguminosa/kounat/internal/tools/config"
)

type PGXConnectionWrapper struct {
	master, slave *pgxpool.Pool
}

func (c *PGXConnectionWrapper) GetMaster(ctx context.Context) (*pgxpool.Pool, error) {
	if c.master == nil {
		return nil, errors.New("connection pool is nil")
	}

	return c.master, nil
}

func (c *PGXConnectionWrapper) GetSlave(ctx context.Context) (*pgxpool.Pool, error) {
	if c.slave == nil {
		return nil, errors.New("connection pool is nil")
	}

	return c.slave, nil
}

func NewPGXClient(ctx context.Context, cfg *config.Config) (tools.PGXClient, error) {
	masterPool, err := connect(ctx, cfg.Database.Master)
	if err != nil {
		return nil, err
	}

	var slavePool *pgxpool.Pool
	slavePool, err = connect(ctx, cfg.Database.Slave)
	if err != nil {
		return nil, err
	}

	return &PGXConnectionWrapper{
		master: masterPool,
		slave:  slavePool,
	}, nil
}

func connect(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	pgxConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	var pool *pgxpool.Pool
	pool, err = pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
