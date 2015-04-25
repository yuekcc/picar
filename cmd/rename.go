/*
命令行参数：重命名照片
使用：picar rename [--prefix <prefix>] [--noarchiving] [--debug] <path=./>
*/

package cmd

import (
	"os"
	"picar/core"

	"github.com/codegangsta/cli"
)

var Rename = cli.Command{
	Name:        "rename",
	Usage:       "rename (and archive) photos",
	Description: "command rename [command options] PATH_OF_PHOTOS",
	Action:      doRename,
	Flags: []cli.Flag{
		cli.StringFlag{"prefix", "", "prefix of new file name", ""},
		//cli.StringFlag{"path", "./", "working folder of photos.", ""},
		cli.BoolFlag{"noarchiving", "no archiving", ""},
		cli.BoolFlag{"debug", "run in debug mode", ""},
	},
}

func doRename(c *cli.Context) {

	path := c.Args().First() // 只处理第一个目录

	// 照片的路径，默认为当前目录
	if path == "" {
		path = "./"
	}

	prefix := c.String("prefix")
	noarchiving := c.Bool("noarchiving")
	debug := c.Bool("debug")

	picar := core.NewPicar(path, prefix, noarchiving, debug)
	err := picar.Parse()

	if err != nil {
		os.Exit(1)
	}
}
