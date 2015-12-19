package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"picar/core"
)

var (
	flagPrefix     string
	flagRenameOnly bool
)

func init() {
	flag.StringVar(&flagPrefix, "prefix", "", "设置文件名的前缀")
	flag.BoolVar(&flagRenameOnly, "renameonly", false, "只重命名文件名")
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		printHelp()
		return
	}

	log.Println("picar, a tool for rename and archiving photos.")
	log.Println("version", _VESION)

	picar(flagPrefix, flagRenameOnly, flag.Args()...)
}

func picar(prefix string, renameOnly bool, path ...string) {
	log.Printf("prefix = %v, renameOnly = %v, path = %v", prefix, renameOnly, path)

	if len(path) == 0 {
		pwd, _ := os.Getwd()
		path = []string{pwd}
	}

	for _, dir := range path {
		log.Println("正在处理目录", dir)
		parser := core.NewParser(prefix, renameOnly, dir)
		err := parser.Parse()
		if err != nil {
			log.Println(err)
		}
	}
}

func printHelp() {
	fmt.Println("使用方法：picar -perfix PREFIX [-renameonly] [path1, path2, ...]")
	flag.PrintDefaults()
}
