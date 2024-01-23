//go:build wireinject
// +build wireinject

package kounatapi

import (
	"context"

	"github.com/google/wire"
	"github.com/leguminosa/kounat/internal/tools/config"
)

func InitServer(
	ctx context.Context,
	cfg *config.Config,
) (Server, error) {
	wire.Build(superSet)
	return &serverImpl{}, nil
}
