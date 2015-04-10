package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	log "github.com/Sirupsen/logrus"
	exif "github.com/rwcarlsen/goexif/exif"
)

type Photo struct {
	Name    string // 原来的文件名（不含拓展名）
	Ext     string // 拓展名
	Path    string // 原始路径（绝对路径，不含文件名）
	Newname string // 新文件名（不含拓展名）
	Archdir string // 归档目录
	Count   int    // 文件计数（应对连拍，部分机器连拍会产生同一个拍摄时间）
	file    string // 全路径（包含文件名、拓展名）
}

func NewPhoto(file string) *Photo {
	// for debug
	//log.Debug("PARSEING FILE: ", file)

	return &Photo{
		Name:  filepath.Base(file),
		Ext:   filepath.Ext(file),
		Path:  filepath.Dir(file),
		Count: 0,
		file:  file,
	}
}
func (self *Photo) GenName(prefix string) error {

	log.Debug("\t- GENERRATE NEW FILE NAME OF FILE: ")
	log.Debug("\t\t| ", self.file)

	jpgfile, err := os.Open(self.file)
	defer jpgfile.Close()

	if err != nil {
		log.Debug("\t-! OS.OPEN ERROR: ")
		log.Debug("\t\t", err)
		return err
	}

	exifdata, err := exif.Decode(jpgfile)
	if err != nil {
		log.Debug("\t-! GOT A EXIF_DECODE ERROR: ")
		log.Debug("\t\t", err)
		return err
	}

	dt, err := exifdata.DateTime()
	if err != nil {
		log.Debug("\t-! GETTING DATETIME ERROR: ")
		log.Debug("\t\t", err)
		return err
	}

	dtsplited := strings.Split(dt.Format("20060102 150405"), " ") // 时间日期格式化~

	// to-do：没有考虑连拍的情况
	if prefix == "" {
		self.Newname = fmt.Sprintf("%s-%s%s", dtsplited[0], dtsplited[1], self.Ext)
	} else {
		self.Newname = fmt.Sprintf("%s-%s-%s%s", prefix, dtsplited[0], dtsplited[1], self.Ext)
	}

	self.Archdir = string(dtsplited[0][:6])
	return nil
}
