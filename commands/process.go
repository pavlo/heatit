package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/pavlo/heatit/app"
	"log"
)

func PerformTheProcessCommand(c *cli.Context) {
	log.Println("I am the Process Command!")

	p, _ := app.NewParameters("foo")
	fmt.Println(p)
}
