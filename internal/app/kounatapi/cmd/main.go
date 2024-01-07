package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/leguminosa/kounat/internal/app/kounatapi"
	"github.com/leguminosa/kounat/internal/tools/config"
)

func Main() {
	e := echo.New()
	cfg := config.New()

	e.Logger.Fatal(kounatapi.NewServer(e, cfg).Start())
}
