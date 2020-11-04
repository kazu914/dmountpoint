package main

import (
	"log"
	"os"

	"github.com/kazu914/dmountpoint/api"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "dmountpoint"
	app.Usage = "dmountpoint shows mount points list of each containers"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "apiVersion, apiversion",
			Value: "default",
			Usage: "API version",
		},
	}

	app.Action = func(c *cli.Context) error {
		api.RunDMountPoint(c.String("apiVersion"))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
