package main

import (
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	cli.VersionFlag = cli.BoolFlag{
		Name:  "print-version, V",
		Usage: "print only the version",
	}

	app := cli.NewApp()
	app.Name = "lit"
	app.Version = "1.0.0"
	app.Run(os.Args)
}
