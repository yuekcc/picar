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
	// FINISH 表示任务已经结束
	FINISH = struct{}{}
)

// Task 表示一个重命名任务
//
type Task struct {
	dir         string
	prefix      string
	renameOnly  bool
	parseVideos bool
}

// CreateTask 创建一个新任务
//
func CreateTask(prefix string, renameOnly bool, videos bool, path string) *Task {
	return &Task{
		dir:         path,
		prefix:      prefix,
		parseVideos: videos,
		renameOnly:  renameOnly,
	}
}

// Execute 执行任务
func (t *Task) Execute() error {
	log.Println("正在处理文件列表")
	files, err := t.getFileList()
	if err != nil {
		return err
	}

	log.Println("过滤照片文件")

	finish := make(chan struct{})
	count := 0
	for i := 0; i < len(files); i++ {
		file := files[i]

		if file == "" {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file))
		switch ext {
		case JPG:
			log.Println("\t- 照片: ", file)
			count++
			go t.parse(file, false, finish)
		case MP4, MOV:
			log.Println("\t- 影片: ", file)
			go t.parse(file, true, finish)
			count++
		default:
			log.Println("\t- 忽略: ", file)
		}
	}

	for i := 0; i < count; i++ {
		<-finish
	}

	log.Println("一共处理了 ", count, " 个文件")
	return nil
}

// 取得文件列表
//
func (t *Task) getFileList() ([]string, error) {
	log.Println("读取目录：", t.dir)

	items, err := os.ReadDir(t.dir)
	if err != nil {
		return nil, err
	}

	itemsLen := len(items)
	result := make([]string, itemsLen)
	for i := 0; i < itemsLen; i++ {
		item := items[i]

		// 忽略子目录
		if item.IsDir() {
			continue
		}

		result[i] = filepath.Join(t.dir, item.Name())
	}

	return result, nil
}

// 重命名（和归档）照片
//
// 流程：
// 1. 如果需要归档照片，则生成新文件名时，加上要放置的目录；否则只重命名文件名
// 2. 将照片重新命名为新文件名
//
func (t *Task) parse(file string, isVideoFile bool, done chan struct{}) {
	onDone := func() {
		done <- FINISH
	}

	onError := func(err interface{}) {
		log.Println(err)

		// 如果出错了，忽略对该文件的处理
		onDone()
	}

	onPanic := func() {
		if err := recover(); err != nil {
			onError(err)
			return
		}
	}

	defer onPanic()

	log.Println("\t- 正在处理：", file)

	newPath := ""
	mf := NewMediaFile(file)

	for counter := 0; ; counter++ {
		//取得新文件名
		err := mf.SetNewFilename(t.prefix, counter, isVideoFile)
		if err != nil {
			onError(err)
			return
		}

		if t.renameOnly {
			// 只修改文件名
			newPath = filepath.Join(mf.Dir, mf.NewFilename)
		} else {
			// 归档照片
			newPath = filepath.Join(mf.Dir, mf.ArchFolder, mf.NewFilename)
			err := os.MkdirAll(filepath.Join(mf.Dir, mf.ArchFolder), 0777)
			if err != nil {
				panic("create directory failed, " + err.Error())
			}
		}

		log.Println("\t- 重命名为：", newPath)

		// 首先检查是否已经储存一样的新文件名的文件
		found, err := isExists(newPath)
		if err != nil {
			onError(err)
			return
		}

		if found {
			continue
		}

		// 没有同名文件开始重命名文件
		os.Rename(file, newPath)
		onDone()
		return
	}
}
