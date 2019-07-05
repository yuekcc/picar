package picarcore

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

func fromExif(path string) ([]string, error) {
	fp, err := os.Open(path)
	defer fp.Close()

	if err != nil {
		return nil, err
	}

	exifData, err := exif.Decode(fp)
	if err != nil {
		return nil, err
	}

	dt, err := exifData.DateTime()
	if err != nil {
		return nil, err
	}

	split := strings.Split(dt.Format("20060102 150405"), " ")
	return split, nil
}

// 处理未知命名格式
//
func fromUnknownNamingFormat(filename string) ([]string, error) {
	ts := time.Now()
	return []string{filename, ts.Format("20060102150405")}, nil
}

// 名字例如：img_20151106_212111
//
func fromFilename(name string) ([]string, error) {
	baseName := filepath.Base(name)
	splits := strings.Split(baseName, ".")
	formatted := strings.ToLower(strings.Join(splits[:len(splits)-1], "."))

	spSet := []string{"_", "-", " "}

	var datePart string
	var timePart string

	expected := false
	for _, sp := range spSet {
		if strings.Contains(formatted, sp) {
			substr := strings.Split(formatted, sp)
			size := len(substr)

			if size == 3 {
				datePart = substr[1]
				timePart = substr[2]
				expected = true
				break
			}

			if size == 2 {
				datePart = substr[0]
				timePart = substr[1]
				expected = true
				break
			}
		}
	}

	if expected {
		return []string{datePart, timePart}, nil
	}

	return fromUnknownNamingFormat(formatted)
}

type Photo struct {
	Dir         string // 当前位置（目录）
	NewFilename string // 新文件名
	ArchFolder  string // 归档目录
	ext         string // 拓展名
	currentPath string // 当前路径（包含文件名）
}

func NewPhoto(path string) *Photo {
	return &Photo{
		ext:         filepath.Ext(path),
		Dir:         filepath.Dir(path),
		ArchFolder:  "other",
		currentPath: path,
	}
}

func (p *Photo) UpdateName(prefix string, counter int) error {
	splits := []string{}
	splits, err := fromExif(p.currentPath)
	if err != nil {
		splits, _ = fromFilename(p.currentPath)
	}

	var buf bytes.Buffer
	if prefix != "" {
		buf.WriteString(prefix)
	}

	buf.WriteString(fmt.Sprintf("%s-%s", splits[0], splits[1]))

	// 处理连拍的情况。连拍时，只会产生一个 Exif 信息（时间是相同的）。
	if counter > 0 {
		buf.WriteString(strconv.Itoa(counter))
	}

	buf.WriteString(p.ext)

	p.NewFilename = buf.String()

	// 如果文件名符合长度，使用文件名的截取部分
	if len(splits[0]) > 6 {
		p.ArchFolder = splits[0][:6]
	}

	return nil
}
