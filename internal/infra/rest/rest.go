package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tesarwijaya/night-owl/internal/config"
	healthz_controller "github.com/tesarwijaya/night-owl/internal/domain/healthz/controller"
	player_controller "github.com/tesarwijaya/night-owl/internal/domain/player/controller"
	team_controller "github.com/tesarwijaya/night-owl/internal/domain/team/controller"
	"go.uber.org/dig"
)

type RestController struct {
	dig.In
	HealthzController healthz_controller.HealthzController
	PlayerController  player_controller.PlayerController
	TeamController    team_controller.TeamController
}

type RestServer struct {
	Server *echo.Echo
	Config *config.Config
}

func NewRestServer(c *config.Config, controllers RestController) RestServer {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	controllers.HealthzController.SetRouter(e)
	controllers.PlayerController.SetRouter(e)
	controllers.TeamController.SetRouter(e)

	return RestServer{
		Server: e,
		Config: c,
	}
}

func (s *RestServer) Start() error {
	return s.Server.Start(fmt.Sprintf(":%s", s.Config.Port))
}
