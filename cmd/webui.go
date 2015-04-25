/*
命令行参数：启动 webui
使用：picar webui --port <ip:port>
*/

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
	fmt.Printf("Start Web UI on http://localhost%s/\n", port)
	web.UI(port)
}
