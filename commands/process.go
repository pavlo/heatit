package commands

import (
	"github.com/codegangsta/cli"
	"github.com/pavlo/heatit/app"
)

func perform(c *cli.Context) {
	engine := app.NewEngine(c)
	engine.Process()
}

func GetProcessCommand() cli.Command {
	return cli.Command{
		Name:   "process",
		Usage:  "Processess a YAML template",
		Action: perform,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "source, s",
				Value: "heat.yaml",
				Usage: "Source HEAT template to process",
			},
			cli.StringFlag{
				Name:  "destination, d",
				Value: "result.yaml",
				Usage: "Destination file where the resulting YAML will be saved",
			},
			cli.StringFlag{
				Name:  "params, p",
				Value: "",
				Usage: "A flat YAML file (k/v) to take parameters from",
			},
			cli.StringSliceFlag{
				Name:  "param-override, P",
				Value: &cli.StringSlice{},
				Usage: "Override parameters loaded from file (--params arg)",
			},
		},
	}
}
