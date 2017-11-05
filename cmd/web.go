package cmd

import (
	"github.com/urfave/cli"
	"io/ioutil"
	"github.com/Dell-/goci/pkg/settings"
	"github.com/Dell-/goci/routers"
	"path"
	log "github.com/go-clog/clog"
	"github.com/mcuadros/go-version"
	"gopkg.in/macaron.v1"
)

var Web = cli.Command{
	Name:        "web",
	Usage:       "Start web server",
	Description: `Goci web server is the only thing you need to run, and it takes care of all the other things for you`,
	Action:      runWeb,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "port, p",
			Value: "3000",
			Usage: "Temporary port number to prevent conflict",
		},
	},
}

// checkVersion checks if binary matches the version of templates files.
func checkVersion() {
	// Client
	fileVer := path.Join(settings.SERVER.StaticRootPath, "/.VERSION")
	data, err := ioutil.ReadFile(fileVer)
	if err != nil {
		log.Fatal(2, "Fail to read '%s': %v", fileVer, err)
	}
	tplVer := string(data)
	if tplVer != settings.APP.AppVer {
		if version.Compare(tplVer, settings.APP.AppVer, ">") {
			log.Fatal(
				2,
				"[%s] Binary version is lower than client file version, did you forget to recompile Gogs?",
				fileVer,
			)
		} else {
			log.Fatal(
				2,
				"[%s] Binary version is higher than client file version, did you forget to update client files?",
				fileVer,
			)
		}
	}
}

// newMacaron initializes Macaron instance
// https://go-macaron.com/docs
func newMacaron() *macaron.Macaron {
	m := macaron.New()
	m.Use(macaron.Static(path.Join(settings.SERVER.StaticRootPath, "dist")))

	return m
}

func runWeb(ctx *cli.Context) error {
	checkVersion()

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
