package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yuekcc/picar/internal/picarcore"
)

var (
	flagPrefix     string
	flagRenameOnly bool
	flagParseVideo bool
	flagVersion    bool
)

func init() {
	flag.StringVar(&flagPrefix, "prefix", "", "设置文件名的前缀")
	flag.BoolVar(&flagRenameOnly, "renameonly", false, "只重命名文件名")
	flag.BoolVar(&flagParseVideo, "video", false, "处理视频文件")
	flag.BoolVar(&flagVersion, "version", false, "显示版本号")
}

func showVersion() {
	fmt.Printf("picar, version %s (%s)\n", VERSION, COMMIT_ID)
	os.Exit(1)
}

func main() {
	flag.Parse()

	if flagVersion {
		showVersion()
	}

	log.Println("picar, a tool for rename and archiving photos.")
	log.Println("version", VERSION)
	log.Println("a tool from yuekcc, build with love.")

	config := picarcore.Config{
		ParseVideo:   flagParseVideo,
		RenameOnly:   flagRenameOnly,
		Prefix:       flagPrefix,
		PhotoFolders: flag.Args(),
	}

	picarcore.Run(config)
}
