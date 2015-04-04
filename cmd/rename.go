package cmd

import (
	"os"
	"picar/core"

	"github.com/codegangsta/cli"
)

var Rename = cli.Command{
	Name:        "rename",
	Usage:       "rename photos.",
	Description: "command rename [command options] PATH_OF_PHOTOS",
	Action:      doRename,
	Flags: []cli.Flag{
		cli.StringFlag{"prefix", "", "prefix of new file name.", ""},
		//cli.StringFlag{"path", "./", "working folder of photos.", ""},
		cli.BoolFlag{"noarchiving", "no archiving.", ""},
		cli.BoolFlag{"debug", "run in debug mode.", ""},
	},
}

func doRename(c *cli.Context) {

	path := c.Args().First()

	if path == "" {
		path = "./"
	}

	prefix := c.String("prefix")
	noarchving := c.Bool("noarchiving")
	debug := c.Bool("debug")

	picar := core.NewPicar(path, prefix, noarchving, debug)
	err := picar.Parse()

	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
