package cmd

import (
	"github.com/urfave/cli"
	"github.com/Dell-/goci/pkg/settings"
	log "github.com/go-clog/clog"
	"github.com/Dell-/goci/pkg/global"
)

var Test = cli.Command{
	Name:        "test",
	Usage:       "Command for test",
	Description: "",
	Action:      runTest,
	Flags: []cli.Flag{
	},
}

func runTest(ctx *cli.Context) error {
	global.Initialize()

	log.Info("APP: %v", settings.APP)
	log.Info("SERVER: %v", settings.SERVER)
	log.Info("DB: %v", settings.DB)

	return nil
}
