package main

import (
	"os"
	"picar/cmd"

	"github.com/codegangsta/cli"
)

const (
	VERSION = "13.2"
)

func main() {
	app := cli.NewApp()
	app.Name = "picar"
	app.Usage = "A photos rename and archive tool"
	app.Author = "Yuekcc <yuekcc@qq.com>"
	app.Version = VERSION
	app.Commands = []cli.Command{
		cmd.Version,
		cmd.Rename, // 命令行
		cmd.Webui,  // web-ui
	}
	app.Run(os.Args)
}
