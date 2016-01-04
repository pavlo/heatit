package main

import (
	"github.com/codegangsta/cli"
	"github.com/pavlo/heatit/commands"
	"os"
)

const (
	VERSION = "0.0.2"
)

func main() {
	app := cli.NewApp()

	app.Name = "heatit"
	app.HelpName = app.Name
	app.Version = VERSION

	app.Usage = "A command line tool that simplifies HEAT templates authoring and processing"
	app.Flags = appFlags()
	app.Commands = appCommands()

	app.Run(os.Args)
}

func appCommands() []cli.Command {
	return []cli.Command{
		commands.GetProcessCommand(),
	}
}

func appFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "show more output",
		},
	}
}
