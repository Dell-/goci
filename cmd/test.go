package cmd

import (
	"github.com/urfave/cli"
	"github.com/Dell-/goci/pkg/settings"
	log "github.com/go-clog/clog"
	"github.com/Dell-/goci/models"
)

var Test = cli.Command{
	Name:        "test",
	Usage:       "Command for test",
	Description: "",
	Action:      runTest,
	Flags: []cli.Flag{
	},
}

func init() {
	settings.NewContext()
}

func runTest(ctx *cli.Context) error {
	log.Info("APP: %v", settings.APP)
	log.Info("SERVER: %v", settings.SERVER)
	log.Info("DB: %v", settings.DB)

	err := models.NewEngine();
	if err != nil {
		log.Fatal(2, "DB fail: %v", err)
	}

	log.Info("DB ping: %v", models.Ping())

	return nil
}
