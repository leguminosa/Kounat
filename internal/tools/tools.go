package tools

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PGXClient interface {
	GetMaster(ctx context.Context) (*pgxpool.Pool, error)
	GetSlave(ctx context.Context) (*pgxpool.Pool, error)
}
