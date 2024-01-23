package kounatapi

import (
	"github.com/labstack/echo/v4"
	"github.com/leguminosa/kounat/internal/tools/config"
)

type (
	Server interface {
		Start() error
	}
	serverImpl struct {
		echo *echo.Echo
		port string
	}
)

func NewServer(
	cfg *config.Config,
	e *echo.Echo,
) Server {
	return &serverImpl{
		echo: e,
		port: cfg.API.Port,
	}
}

func (s *serverImpl) Start() error {
	s.echo.Logger.Info("kounatapi running on port", s.port)
	return s.echo.Start(s.port)
}
