package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var version string = "dev"

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print only the version",
	}

	app := &cli.App{
		Name:        "aws-tools",
		Usage:       "small tools around aws api",
		Version:     version,
		Commands:    []*cli.Command{NewAllowIpCommand()},
		HideVersion: false,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
