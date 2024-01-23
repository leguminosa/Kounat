package tools

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type PGXPool interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}
