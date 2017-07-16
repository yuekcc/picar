package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"picar/core"
)

var (
	flagPrefix      string
	flagNoArchiving bool
	flagParseVideos bool
)

type PicarConfig struct {
	ParseVideos    bool
	NoArchiving    bool
	Prefix         string
	PicruteFolders []string
}

func init() {
	flag.StringVar(&flagPrefix, "prefix", "", "设置文件名的前缀")
	flag.BoolVar(&flagNoArchiving, "noarchiving", false, "不归档文件")
	flag.BoolVar(&flagParseVideos, "videos", false, "处理视频文件")
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		printHelp()
		os.Exit(-1)
		return
	}

	log.Println("picar, a tool for rename and archiving photos.")
	log.Println("version", _VERSION)
	log.Println("a tool from yuekcc, build with love.")

	config := PicarConfig{
		ParseVideos:    flagParseVideos,
		NoArchiving:    flagNoArchiving,
		Prefix:         flagPrefix,
		PicruteFolders: flag.Args(),
	}

	picar(config)
}

func picar(config PicarConfig) {
	log.Printf("prefix = %v, noArchiving = %v, path = %v", config.Prefix, config.NoArchiving, config.PicruteFolders)

	var dirs []string
	if len(config.PicruteFolders) == 0 {
		pwd, _ := os.Getwd()
		dirs = []string{pwd}
	}

	for _, dir := range dirs {
		log.Println("正在处理目录", dir)
		parser := core.NewParser(config.Prefix, config.NoArchiving, config.ParseVideos, dir)
		err := parser.Parse()
		if err != nil {
			log.Println(err)
		}
	}
}

func printHelp() {
	fmt.Println("使用方法：picar -perfix PREFIX [-noarchiving] [-videos] [path1, path2, ...]")
	flag.PrintDefaults()
}
