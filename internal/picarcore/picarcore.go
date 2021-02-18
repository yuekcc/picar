package picarcore

import (
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	ParseVideo   bool
	RenameOnly   bool
	Prefix       string
	PhotoFolders []string
}

func Run(config Config) {
	log.Printf("prefix = %v, noArchiving = %v, path = %v", config.Prefix, config.RenameOnly, config.PhotoFolders)

	pwd, _ := os.Getwd()
	dirs := config.PhotoFolders
	if len(dirs) == 0 {
		dirs = append(dirs, pwd)
	}

	for _, dir := range dirs {
		targetDir := dir
		log.Println("正在处理目录", targetDir)

		if !filepath.IsAbs(targetDir) {
			targetDir = filepath.Join(pwd, targetDir)
		}

		task := CreateTask(config.Prefix, config.RenameOnly, config.ParseVideo, targetDir)
		err := task.Execute()
		if err != nil {
			log.Println(err)
		}
	}
}
