package picarcore

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	JPG = ".jpg"
	MOV = ".mov"
	MP4 = ".mp4"
)

var (
	FINISH = struct{}{}
)

// 重命名目录任务
//
type Task struct {
	dir         string
	prefix      string
	noArchiving bool
	parseVideos bool
	files       []string
}

func CreateTask(prefix string, noArchiving bool, videos bool, path string) *Task {
	return &Task{
		dir:         path,
		prefix:      prefix,
		parseVideos: videos,
		noArchiving: noArchiving,
	}
}

func (t *Task) Run() error {
	log.Println("正在处理文件列表")
	err := t.getFileList()
	if err != nil {
		return err
	}

	finish := make(chan struct{})

	index := 0
	log.Println("过滤照片文件")
	for _, file := range t.files {
		ext := strings.ToLower(filepath.Ext(file))
		switch ext {
		case JPG:
			log.Println("\t- 照片: ", file)
			index++
			go t.parse(file, false, finish)
		case MP4, MOV:
			log.Println("\t- 影片: ", file)
			go t.parse(file, true, finish)
		default:
			log.Println("\t- 忽略: ", file)
		}
	}

	for i := 0; i < index; i++ {
		<-finish
	}

	log.Println("一共处理了 ", index, " 个文件")
	return nil
}

// 取得文件列表
//
func (t *Task) getFileList() (err error) {

	log.Println("读取目录：", t.dir)

	items, err := os.ReadDir(t.dir)
	if err != nil {
		return err
	}

	for _, item := range items {
		// 忽略子目录
		if item.IsDir() {
			continue
		}

		file := filepath.Join(t.dir, item.Name())

		t.files = append(t.files, file)
	}
	return nil
}

// 重命名（和归档）照片
// 原理：
// 如果需要归档照片，则生成新文件名时，加上要放置的目录
// 然后，将照片重新命名为新文件名
//
func (t *Task) parse(file string, isVideoFile bool, done chan struct{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			done <- FINISH
			return
		}
	}()

	log.Println("\t- 正在处理：", file)

	newPath := ""
	photo := NewPhoto(file)

	for counter := 0; ; counter++ {
		//取得新文件名
		err := photo.UpdateName(t.prefix, counter, isVideoFile)
		if err != nil {
			done <- FINISH
			return
		}

		if t.noArchiving {
			// 不归档照片
			newPath = filepath.Join(photo.Dir, photo.NewFilename)
		} else {
			// 归档照片
			newPath = filepath.Join(photo.Dir, photo.ArchFolder, photo.NewFilename)
			err := os.MkdirAll(filepath.Join(photo.Dir, photo.ArchFolder), 0777)
			if err != nil {
				panic("create directory failed, " + err.Error())
			}
		}

		log.Println("\t- 重命名为：", newPath)

		// 首先检查是否已经储存一样的新文件名的文件
		found, err := isExists(newPath)
		if err != nil {
			done <- FINISH
			return
		}

		if found {
			continue
		}

		// 没有同名文件开始重命名文件
		err = os.Rename(file, newPath)
		if err != nil {
			done <- FINISH
			return
		}

	}
}
