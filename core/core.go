package core

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Picar struct {
	path        string
	prefix      string
	noArchiving bool
	parseVideos bool
	filelist    []string
}

func NewParser(prefix string, noArchiving bool, videos bool, path string) *Picar {
	return &Picar{
		path:        path,
		prefix:      prefix,
		parseVideos: videos,
		noArchiving: noArchiving,
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
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
		ext := strings.ToLower(filepath.Ext(file))
		switch ext {
		case ".jpg":
			log.Println("\t- 照片: ", file)
			index++
			go self.parseImage(file, ch)
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
func (self *Picar) parseImage(file string, done chan bool) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			done <- true
			return
		}
	}()

	log.Println("\t- 正在处理：", file)

	newfullpath := ""
	photo := NewPhoto(file)

	for counter := 0; ; counter++ {
		//取得新文件名
		err := photo.GenName(self.prefix, counter)
		if err != nil {
			done <- true
			return
		}

		if self.noArchiving {
			// 不归档照片
			newfullpath = filepath.Join(photo.Path, photo.NewFilename)
		} else {
			// 归档照片
			newfullpath = filepath.Join(photo.Path, photo.ArchFolder, photo.NewFilename)
			os.MkdirAll(filepath.Join(photo.Path, photo.ArchFolder), 0777)
		}

		log.Println("\t- 重命名为：", newfullpath)

		// 首先检查是否已经储存一样的新文件名的文件
		found, err := exists(newfullpath)
		if err != nil {
			done <- true
			return
		}

		if found {
			continue
		} else {
			err = os.Rename(file, newfullpath)
			if err != nil {
				done <- true
				return
			}
		}
	}
}
