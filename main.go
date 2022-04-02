package main

import (
	"github.com/tesarwijaya/night-owl/internal/config"
	"github.com/tesarwijaya/night-owl/internal/databases/sql"
	healthz_controller "github.com/tesarwijaya/night-owl/internal/domain/healthz/controller"
	player_controller "github.com/tesarwijaya/night-owl/internal/domain/player/controller"
	player_repository "github.com/tesarwijaya/night-owl/internal/domain/player/repository"
	player_service "github.com/tesarwijaya/night-owl/internal/domain/player/service"
	team_controller "github.com/tesarwijaya/night-owl/internal/domain/team/controller"
	team_repository "github.com/tesarwijaya/night-owl/internal/domain/team/repository"
	team_service "github.com/tesarwijaya/night-owl/internal/domain/team/service"
	"github.com/tesarwijaya/night-owl/internal/infra/rest"
	"go.uber.org/dig"
)

func main() {
	c := dig.New()

	c.Provide(rest.NewRestServer)
	c.Provide(config.NewConfig)
	c.Provide(sql.NewSqlDB)

	c.Provide(healthz_controller.NewHealthzController)

	c.Provide(player_controller.NewPlayerController)
	c.Provide(player_service.NewPlayerService)
	c.Provide(player_repository.NewPlayerReposity)
	c.Provide(player_repository.NewPlayerData)

	c.Provide(team_controller.NewTeamController)
	c.Provide(team_service.NewTeamService)
	c.Provide(team_repository.NewTeamReposity)
	c.Provide(team_repository.NewTeamData)

	err := c.Invoke(func(server rest.RestServer) {
		server.Start()
	})

	if err != nil {
		panic(err)
	}
}
