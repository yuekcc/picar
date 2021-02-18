package main

import (
	"flag"
	"log"
	"os"

	"picar/internal/picarcore"
)

var (
	flagPrefix     string
	flagRenameOnly bool
	flagParseVideo bool
)

func init() {
	flag.StringVar(&flagPrefix, "prefix", "", "设置文件名的前缀")
	flag.BoolVar(&flagRenameOnly, "renameonly", false, "只重命名文件名")
	flag.BoolVar(&flagParseVideo, "video", false, "处理视频文件")
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(-1)
		return
	}

	log.Println("picar, a tool for rename and archiving photos.")
	log.Println("version", _VERSION)
	log.Println("a tool from yuekcc, build with love.")

	config := picarcore.Config{
		ParseVideo:   flagParseVideo,
		RenameOnly:   flagRenameOnly,
		Prefix:       flagPrefix,
		PhotoFolders: flag.Args(),
	}

	picarcore.Run(config)
}
