package commands

import (
	"github.com/codegangsta/cli"
	"log"
)

func PerformTheProcessCommand(c *cli.Context) {
	log.Println("I am the Process Command!")
}