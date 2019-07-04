package main

import (
	"flag"
	"log"
	"os"

	"picar/internal/picarcore"
)

var (
	flagPrefix      string
	flagNoArchiving bool
	flagParseVideo  bool
)

func init() {
	flag.StringVar(&flagPrefix, "prefix", "", "设置文件名的前缀")
	flag.BoolVar(&flagNoArchiving, "noarchiving", false, "不归档文件")
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
		NoArchiving:  flagNoArchiving,
		Prefix:       flagPrefix,
		PhotoFolders: flag.Args(),
	}

	picarcore.Run(config)
}
