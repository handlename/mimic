package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

// Version of this application.
var Version = "HEAD"

func main() {
	app := cli.NewApp()
	app.Name = "mimic"
	app.Usage = "Simple http mock server"
	app.Version = Version
	app.Author = "handlename"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: "config.json",
			Usage: "path to config file",
		},
	}
	app.Action = func(c *cli.Context) {
		conf, err := NewConfig(c.String("config"))

		if err != nil {
			log.Fatalf("failed to load config: %s", err)
		}

		mimic := NewMimic(&conf)
		mimic.Start()
	}
	app.Run(os.Args)
}