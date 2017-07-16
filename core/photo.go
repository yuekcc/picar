package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"bytes"

	"strconv"

	exif "github.com/rwcarlsen/goexif/exif"
)

type Photo struct {
	Name        string // 原来的文件名（不含拓展名）
	Ext         string // 拓展名
	Path        string // 原始路径（绝对路径，不含文件名）
	NewFilename string // 新文件名（不含拓展名）
	ArchFolder  string // 归档目录
	Count       int    // 文件计数（应对连拍，部分机器连拍会产生同一个拍摄时间）
	fullpath    string // 全路径（包含文件名、拓展名）
}

func NewPhoto(file string) *Photo {
	return &Photo{
		Name:     filepath.Base(file),
		Ext:      filepath.Ext(file),
		Path:     filepath.Dir(file),
		Count:    0,
		fullpath: file,
	}
}

func (self *Photo) GenName(prefix string, counter int) error {
	var dtsplited []string
	dtsplited, err := genNameFromExif(self.fullpath)
	if err != nil {
		dtsplited, _ = getNameFromFilename(self.fullpath)
	}

	var buf bytes.Buffer
	if prefix != "" {
		buf.WriteString(prefix)
	}

	buf.WriteString(fmt.Sprintf("%s-%s", dtsplited[0], dtsplited[1]))

	// 处理连拍的情况。连拍时，只会产生一个 Exif 信息（时间是相同的）。
	if counter > 0 {
		buf.WriteString(strconv.Itoa(counter))
	}

	buf.WriteString(self.Ext)

	self.NewFilename = buf.String()
	self.ArchFolder = string(dtsplited[0][:6])
	return nil
}

func genNameFromExif(jpeg string) ([]string, error) {
	fp, err := os.Open(jpeg)
	defer fp.Close()

	if err != nil {
		return nil, err
	}

	exifdata, err := exif.Decode(fp)
	if err != nil {
		return nil, err
	}

	dt, err := exifdata.DateTime()
	if err != nil {
		return nil, err
	}

	dtsplited := strings.Split(dt.Format("20060102 150405"), " ")
	return dtsplited, nil
}

// 名字例如：IMG_20151106_212111
//
func getNameFromFilename(jpeg string) ([]string, error) {
	nameStr := filepath.Base(jpeg)
	name := strings.Split(nameStr, ".")
	str := strings.ToUpper(name[0])

	spSet := []string{"_", "-", " "}

	var dateStr string
	var timeStr string

	for _, sp := range spSet {
		if strings.Contains(str, sp) {
			substr := strings.Split(str, sp)
			size := len(substr)
			if size == 3 {
				dateStr = substr[1]
				timeStr = substr[2]
			} else {
				dateStr = substr[0]
				timeStr = substr[1]
			}
			break
		}
	}
	return []string{dateStr, timeStr}, nil
}
