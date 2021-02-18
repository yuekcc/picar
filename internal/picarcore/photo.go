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

func genFromExif(path string) ([]string, error) {
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
func genFromUnknownNamingFormat(filename string) ([]string, error) {
	ts := time.Now()
	return []string{filename, ts.Format("20060102150405")}, nil
}

// 名字例如：img_20151106_212111
//
func genFromFilename(name string) ([]string, error) {
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

	return genFromUnknownNamingFormat(formatted)
}

func genFilename(path string, isVideoFile bool) ([]string, error) {
	if isVideoFile {
		splits, _ := genFromFilename(path)
		return splits, nil
	}

	splits := []string{}
	splits, err := genFromExif(path)
	if err != nil {
		splits, _ = genFromFilename(path)
	}

	return splits, nil
}

// MediaFile 表示一个照片/视频文件
//
type MediaFile struct {
	Dir         string // 当前位置（目录）
	NewFilename string // 新文件名
	ArchFolder  string // 归档目录
	ext         string // 拓展名
	currentPath string // 当前路径（包含文件名）
}

// NewMediaFile 创建一个 MediaFile 结构
//
func NewMediaFile(path string) *MediaFile {
	return &MediaFile{
		ext:         filepath.Ext(path),
		Dir:         filepath.Dir(path),
		ArchFolder:  "other",
		currentPath: path,
	}
}

// SetNewFilename 设置新的文件名
//
func (p *MediaFile) SetNewFilename(prefix string, counter int, isVideoFile bool) error {
	splits, err := genFilename(p.currentPath, isVideoFile)
	if err != nil {
		return err
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

	if isVideoFile {
		p.ArchFolder = "video"
	}

	// 如果文件名符合长度，使用文件名的截取部分
	if len(splits[0]) > 6 && !isVideoFile {
		p.ArchFolder = splits[0][:6]
	}

	return nil
}
