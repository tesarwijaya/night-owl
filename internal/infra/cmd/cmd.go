package cmd

import (
	"github.com/tesarwijaya/night-owl/internal/config"
	"github.com/tesarwijaya/night-owl/internal/domain/system_tools/migration/service"
	"github.com/tesarwijaya/night-owl/internal/infra/rest"
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

type CmdService struct {
	dig.In
	C          config.Config
	RestServer rest.RestServer
	Migration  service.MigrationService
}

func NewCmd(cmd CmdService) *cli.App {
	return &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "run night-owl server",
				Action: func(c *cli.Context) error {
					return cmd.RestServer.Start()
				},
			},
			{
				Name:    "migrate",
				Aliases: []string{"m"},
				Usage:   "migrate database",
				Subcommands: []*cli.Command{
					{
						Name:  "up",
						Usage: "run migration up",
						Action: func(c *cli.Context) error {
							return cmd.Migration.Up()
						},
					},
					{
						Name:  "down",
						Usage: "run migration down",
						Action: func(c *cli.Context) error {
							return cmd.Migration.Down()
						},
					},
				},
			},
		},
	}
}
