package tools

import (
	"context"
)

type PGXClient interface {
	GetMaster(ctx context.Context) (PGXPool, error)
	GetSlave(ctx context.Context) (PGXPool, error)
}
