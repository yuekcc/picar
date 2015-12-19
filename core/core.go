package core

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Picar struct {
	path       string
	prefix     string
	renameOnly bool
	filelist   []string
}

func NewParser(prefix string, renameOnly bool, path string) *Picar {
	return &Picar{
		path:       path,
		prefix:     prefix,
		renameOnly: renameOnly,
	}
}

func (self *Picar) Parse() error {
	log.Println("正在处理文件列表")
	err := self.getFileList()
	if err != nil {
		return err
	}

	ch := make(chan bool)

	index := 0
	log.Println("过滤照片文件")
	for _, file := range self.filelist {
		ext := filepath.Ext(file)
		switch ext {
		case ".jpg":
			log.Println("\t- 照片: ", file)
			index++
			go self.do(file, ch)
		case ".mp4", ".mov":
			log.Println("\t- 影片: ", file)
		default:
			log.Println("\t- 忽略: ", file)
		}
	}

	for i := 0; i < index; i++ {
		<-ch
	}

	log.Println("一共处理了 ", index, " 个文件")
	return nil
}

// 取得文件列表
//
func (self *Picar) getFileList() (err error) {

	log.Println("读取目录：", self.path)

	items, err := ioutil.ReadDir(self.path)
	if err != nil {
		return err
	}

	for _, item := range items {

		// 忽略子目录
		if item.IsDir() {
			continue
		}

		file := filepath.Join(self.path, item.Name())

		self.filelist = append(self.filelist, file)
	}
	return nil
}

// 重命名（和归档）照片
// 原理：
// 如果需要归档照片，则生成新文件名时，加上要放置的目录
// 然后，将照片重新命名为新文件名
//
func (self *Picar) do(file string, done chan bool) {
	log.Println("\t- 正在处理：", file)

	newfullpath := ""
	photo := NewPhoto(file)

	err := photo.GenName(self.prefix) //取得新文件名
	if err != nil {
		done <- true
		return
	}

	// 如果使用了 noarch 标记
	if self.renameOnly {
		// 不归档照片
		newfullpath = filepath.Join(photo.Path, photo.NewFilename)
	} else {
		// 归档照片
		newfullpath = filepath.Join(photo.Path, photo.ArchFolder, photo.NewFilename)
		os.MkdirAll(filepath.Join(photo.Path, photo.ArchFolder), 0777)
	}

	log.Println("\t- 重命名为：", newfullpath)

	// 重命名照片
	// 相当于 shell mv src dest
	err = os.Rename(file, newfullpath)
	if err != nil {
		done <- true
		return
	}

	done <- true
}
