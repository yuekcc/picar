package main

import (
	"flag"
	"log"
	"os"
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

	log.Println("picar, a tool for rename and archiving photos.")
	log.Println("version", _VESION)

	picar(flagPrefix, flagRenameOnly, flag.Args()...)
}

func picar(prefix string, renameOnly bool, path ...string) {
	log.Printf("%v, %v, %v", prefix, renameOnly, path)

	if len(path) == 0 {
		pwd, _ := os.Getwd()
		path = []string{pwd}
	}

	for _, dir := range path {
		log.Println(dir)
	}

}
