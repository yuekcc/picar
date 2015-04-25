/*
命令行参数：显示版本号
使用：picar version
*/

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
