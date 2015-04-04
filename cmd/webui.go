package cmd

import "github.com/codegangsta/cli"

var Webui = cli.Command{
	Name:   "webui",
	Usage:  "run web ui",
	Action: doWebui,
}

func doWebui(c *cli.Context) {
	println("TO BE DONE.")
}
