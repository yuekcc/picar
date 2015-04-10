package core

import (
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
)

type Picar struct {
	path     string
	prefix   string
	noarch   bool
	filelist []string
}

func NewPicar(path string, prefix string, noarch bool, debug bool) *Picar {

	// 进入 debug 模式
	if debug {
		log.SetLevel(log.DebugLevel)
	}

	log.Info("GOT PARAMETERS:")
	log.Info("\t- PATH = ", path)
	log.Info("\t- PREFIX = ", prefix)
	log.Info("\t- NO ARCHIVING = ", noarch)
	log.Info("\t- DEBUG MODE = ", debug)

	return &Picar{
		path:   path,
		prefix: prefix,
		noarch: noarch,
	}
}

func (self *Picar) Parse() error {

	log.Debug("STAGE 1 STARTED.")

	err := self.getFileList()
	if err != nil {
		return err
	}

	log.Debug("STAGE 1 DONE.")

	ch := make(chan int)

	index := 0
	log.Debug("STAGE 2 STARTED.")
	for _, file := range self.filelist {
		ext := filepath.Ext(file)
		switch ext {
		case ".jpg":
			log.Debug("\t- GET A JPG: ", file)
			index++
			//abspath, _ := filepath.Abs(file)
			go self.do(file, ch)
		case ".mp4":
			log.Debug("\t- GET A MOV: ", file)
			//log.Debug("\t- DO NOTHING.")
		default:
			log.Debug("\t- IGNORE FILE: ", file)
		}
	}

	for i := 0; i < index; i++ {
		<-ch
	}
	log.Debug("PARSED ", index, " FILES.")
	log.Debug("STAGE 2 DONE.")
	log.Info("DONE.")
	return nil
}

// 取得文件列表
// fig.1
func (self *Picar) getFileList() (err error) {

	log.Debug("READING DIR: ", self.path)

	items, err := ioutil.ReadDir(self.path)
	if err != nil {
		return err
	}

	for _, item := range items {
		log.Debug("\t- GET A ITEM: ", item.Name())

		// 忽略子目录
		if item.IsDir() {
			//log.Debug("\t\t", item.Name(), "IS DIR! PASSED.")
			continue
		}

		file := filepath.Join(self.path, item.Name())

		//log.Debug("\t\tGET A FILE: ", file)

		self.filelist = append(self.filelist, file)
	}
	return nil // 操作成功就没有 err 了，err = nil。!!!-_-
}

// fig.2
func (self *Picar) do(file string, ch chan int) {
	log.Debug("STAGE 2/RENAME STARTED.")
	log.Debug("\t- PARSING FILE: ", file)

	newfullpath := ""
	photo := NewPhoto(file)
	err := photo.GenName(self.prefix) //取得新文件名
	if err != nil {
		ch <- 1
		return
	}

	if self.noarch {
		newfullpath = filepath.Join(photo.Path, photo.Newname)
	} else {
		newfullpath = filepath.Join(photo.Path, photo.Archdir, photo.Newname)
		os.MkdirAll(filepath.Join(photo.Path, photo.Archdir), 0777)
	}

	log.Debug("\t- SET TO: ", newfullpath)

	err = os.Rename(file, newfullpath)
	if err != nil {
		ch <- 1
		return
	}

	log.Debug("STAGE 2/RENAME DONE.")
	ch <- 1
}
