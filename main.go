package main

import (
	"context"
	"time"

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
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			rest.NewRestServer,
			config.NewConfig,
			sql.NewSqlDB,
			healthz_controller.NewHealthzController,
			player_controller.NewPlayerController,
			player_service.NewPlayerService,
			player_repository.NewPlayerReposity,
			player_repository.NewPlayerData,
			team_controller.NewTeamController,
			team_service.NewTeamService,
			team_repository.NewTeamReposity,
			team_repository.NewTeamData,
		),
		fx.Invoke(func(server rest.RestServer) {
			server.Start()
		}),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		panic(err)
	}

	<-app.Done()
}
