package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "mimic"
	app.Usage = "Simple http mock server"
	app.Version = "v0.1.0"
	app.Author = "handlename"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: "config.json",
			Usage: "path to config file",
		},
		cli.IntFlag{
			Name:  "port",
			Value: 3390,
			Usage: "port of mimic server",
		},
	}
	app.Action = func(c *cli.Context) {
		conf, err := NewConfig(c.String("config"))

		if err != nil {
			log.Fatalf("failed to load config: %s", err)
		}

		mimic := NewMimic(c.Int("port"), &conf)
		mimic.Start()
	}
	app.Run(os.Args)
}
