package cmd

import "github.com/codegangsta/cli"

var Version = cli.Command{
	Name:   "version",
	Usage:  "print version",
	Action: doVersion,
}

func doVersion(c *cli.Context) {
	cli.ShowVersion(c)
}
