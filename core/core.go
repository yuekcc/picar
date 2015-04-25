package core

import (
	"io"
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
	debug    bool
}

func NewPicar(path string, prefix string, noarch bool, debug bool) *Picar {

	// 进入 debug 模式
	if debug {
		log.SetLevel(log.DebugLevel)
	}

	log.Info("参数列表：")
	log.Info("\t- PATH = ", path)
	log.Info("\t- PREFIX = ", prefix)
	log.Info("\t- NO ARCHIVING = ", noarch)
	log.Info("\t- DEBUG MODE = ", debug)

	return &Picar{
		path:   path,
		prefix: prefix,
		noarch: noarch,
		debug:  debug,
	}
}

func (self *Picar) SetOutput(out io.Writer) {
	log.SetOutput(out)
	log.SetFormatter(new(logFormatter))
}

func (self *Picar) Parse() error {

	log.Debug("STAGE 1 获取文件列表")

	err := self.getFileList()
	if err != nil {
		return err
	}

	log.Debug("STAGE 1 完成")

	ch := make(chan int)

	index := 0
	log.Debug("STAGE 2 过滤照片文件")
	for _, file := range self.filelist {
		ext := filepath.Ext(file)
		switch ext {
		case ".jpg":
			log.Debug("\t- 照片: ", file)
			index++
			//abspath, _ := filepath.Abs(file)
			go self.do(file, ch)
		case ".mp4":
			log.Debug("\t- 影片: ", file)
			//log.Debug("\t- DO NOTHING.")
		default:
			log.Debug("\t- 忽略: ", file)
		}
	}

	for i := 0; i < index; i++ {
		<-ch
	}
	log.Debug("一共处理了 ", index, " 个文件")
	log.Debug("STAGE 2 完成")
	log.Info("完成")
	return nil
}

// 取得文件列表
//
func (self *Picar) getFileList() (err error) {

	log.Debug("读取目录：", self.path)

	items, err := ioutil.ReadDir(self.path)
	if err != nil {
		return err
	}

	for _, item := range items {
		//log.Debug("\t- GET A ITEM: ", item.Name())

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

// 重命名（和归档）照片
// 原理：
//	如果需要归档照片，则生成新文件名时，加上要放置的目录
//	然后，将照片重新命名为新文件名
//
func (self *Picar) do(file string, ch chan int) {
	log.Debug("\t- 正在处理：", file)

	newfullpath := ""
	photo := NewPhoto(file)
	err := photo.GenName(self.prefix) //取得新文件名
	if err != nil {
		ch <- 1
		return
	}

	// 如果使用了 noarch 标记
	if self.noarch {
		// 不归档照片
		newfullpath = filepath.Join(photo.Path, photo.Newname)
	} else {
		// 归档照片
		newfullpath = filepath.Join(photo.Path, photo.Archdir, photo.Newname)
		os.MkdirAll(filepath.Join(photo.Path, photo.Archdir), 0777)
	}

	log.Debug("\t- 重命名为：", newfullpath)

	// 重命名照片
	// 相当于 shell mv src dest
	err = os.Rename(file, newfullpath)
	if err != nil {
		ch <- 1
		return
	}

	ch <- 1
}
