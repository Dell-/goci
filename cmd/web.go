package cmd

import (
	"github.com/urfave/cli"
	"github.com/Dell-/goci/pkg/settings"
	"github.com/Dell-/goci/routers"
	"path"
	"gopkg.in/macaron.v1"
	"github.com/Dell-/goci/pkg/global"
)

var Web = cli.Command{
	Name:        "web",
	Usage:       "Start web server",
	Description: `Goci web server is the only thing you need to run, and it takes care of all the other things for you`,
	Action:      runWeb,
}

// newMacaron initializes Macaron instance
// https://go-macaron.com/docs
func newMacaron() *macaron.Macaron {
	m := macaron.New()
	m.Use(macaron.Static(path.Join(settings.SERVER.StaticRootPath, "dist")))

	return m
}

func runWeb(ctx *cli.Context) error {
	global.Initialize()

	m := newMacaron()

	m.Get("/", routers.Index)

	// TODO: this route must be removed
	m.Get("*", func(ctx *macaron.Context) {
		if ctx.Req.RequestURI != "/" {
			ctx.Redirect("/")
		}
	})
	m.Run(settings.SERVER.HTTPAddr, settings.SERVER.HTTPPort)

	return nil
}
