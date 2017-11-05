/**
 * Copyright 2014 The Andrii Hurzhii Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */
package main

import (
	"github.com/Dell-/goci/cmd"
	"github.com/urfave/cli"
	"os"
	"github.com/Dell-/goci/pkg/settings"
)

func main() {
	app := cli.NewApp()
	app.Name = "Goci"
	app.Usage = "A painless self-hosted CI service"
	app.Version = settings.APP.AppVer
	app.Commands = []cli.Command{
		cmd.Rest,
		cmd.Web,
		cmd.Test,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)

	app.Run(os.Args)
}
