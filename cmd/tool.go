package cmd

import (
	"github.com/urfave/cli"
	"github.com/Dell-/goci/pkg/global"
	"fmt"
	"github.com/Dell-/goci/models"
	"github.com/icrowley/fake"
	log "github.com/go-clog/clog"
)

var Tool = cli.Command{
	Name:        "tool",
	Usage:       "Tool for development",
	Description: "",
	Action:      runTool,
	Subcommands: []cli.Command{
		{
			Name:  "create",
			Usage: "create new user",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "email, e",
					Usage: "Use --email or -e for set User Email",
				},
				cli.StringFlag{
					Name:  "password, p",
					Usage: "Use --password or -p for set User Password",
				},
				cli.BoolFlag{
					Name:  "with-token, t",
					Usage: "Use --with-token or -t for set User Access Token",
				},
			},
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					return nil
				}

				commandInitialize()

				switch  c.Args().Get(0) {
				case "user":
					email := c.String("email")
					password := c.String("password")
					withToken := c.Bool("with-token")

					if len(email) == 0 || len(password) == 0 {
						fmt.Println("Email and Password is required fields.")
						return nil
					}

					fmt.Println("Create User entity")
					fmt.Println("--email ", email)
					fmt.Println("--password ", password)

					user := &models.User{
						Email:    email,
						Password: password,
						Username: fake.UserName(),
						FullName: fake.FullName(),
						IsActive: true,
					}

					if err := models.CreateUser(user); err != nil {
						log.Fatal(0, "Cannot create user. %s", err)
					}

					if !withToken {
						return nil
					}
					accessToken := &models.AccessToken{
						UID: user.ID,
					}

					return models.NewAccessToken(accessToken)
				}

				return nil
			},
		},
	},
}

func commandInitialize() {
	global.Initialize()
}

func runTool(ctx *cli.Context) error {
	commandInitialize()

	return nil
}
