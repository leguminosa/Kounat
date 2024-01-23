package kounatapi

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/leguminosa/kounat/internal/app/kounatapi/handler"
	characterModule "github.com/leguminosa/kounat/internal/module/character"
	characterRepo "github.com/leguminosa/kounat/internal/repository/character"
	"github.com/leguminosa/kounat/internal/storage/database/postgres"
	"github.com/leguminosa/kounat/internal/tools/database"
)

var (
	// internal/tools
	toolsSet = wire.NewSet(
		database.NewPGXClient,
	)

	// internal/storage
	storageDatabaseSet = wire.NewSet(
		postgres.NewCharacterDB,
	)

	// internal/repository
	repositorySet = wire.NewSet(
		characterRepo.NewRepository,
	)

	// internal/module
	moduleSet = wire.NewSet(
		characterModule.NewModule,
	)

	// internal/app/kounatapi/handler
	handlerSet = wire.NewSet(
		handler.NewCharacter,
	)

	// aggregate all wire sets here
	superSet = wire.NewSet(
		NewServer,
		provideRouter,
		handlerSet,
		moduleSet,
		repositorySet,
		storageDatabaseSet,
		toolsSet,
	)
)

func provideRouter(
	characterHandler *handler.CharacterHandler,
) *echo.Echo {
	e := echo.New()

	// register your routers here
	e.GET("/characters/:id", characterHandler.GetByID)

	return e
}
