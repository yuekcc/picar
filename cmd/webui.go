package cmd

import (
	"fmt"
	"picar/web"

	"github.com/codegangsta/cli"
)

var Webui = cli.Command{
	Name:   "webui",
	Usage:  "run web ui",
	Action: doWebui,
	Flags: []cli.Flag{
		cli.StringFlag{"port", ":8088", "ip:port", ""},
	},
}

func doWebui(c *cli.Context) {
	port := c.String("port")
	fmt.Printf("Start Server on http://localhost%s/\n", port)
	web.UI(port)
}
