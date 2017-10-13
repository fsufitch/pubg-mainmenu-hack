package server

import (
	"fmt"
	"log"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

// RunCommand runs the command line interface for pubg-mainmenu-hack
func RunCommand() {
	app := cli.NewApp()
	app.Name = "pubg-mainmenu-hack"
	app.Usage = "command line interface to pubg-mainmenu-hack server and utilities"
	app.Commands = []cli.Command{
		{
			Name:      "serve",
			Usage:     "start the pubg-mainmenu-hack web server",
			ArgsUsage: "STATIC_DIR",
			Action: func(c *cli.Context) error {
				staticDir := c.Args().First()
				server, err := newWebServer(staticDir)
				if err != nil {
					log.Fatal(err)
				}

				log.Fatal(server.Start())
				return nil
			},
		},

		{
			Name:   "migrate",
			Usage:  "run all available migrations",
			Action: func(*cli.Context) { fmt.Println("there is no database in use right now") },
		},
	}

	app.Run(os.Args)
}
