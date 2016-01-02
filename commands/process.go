package commands

import (
	"github.com/codegangsta/cli"
	"log"
	"github.com/pavlo/heatit/app"
	"fmt"
)

func PerformTheProcessCommand(c *cli.Context) {
	log.Println("I am the Process Command!")

	p, _ := app.NewParameters("foo")
	fmt.Println(p)
}