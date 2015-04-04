// PICAR：一个照片归档工具。
// 将照片按201501命名的目录归档照片
// 照片命名格式：[prefix-]20150102-150406.jpg
// by yuekcc@qq.com
// 版本：10
// 更新日期：2015-01-13 1121PM

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
	//"github.com/siddontang/go/log"
	log "github.com/Sirupsen/logrus"
)

// version
const VERSION = "11"

type Pic struct {
	name    string // 原来的文件名
	ext     string // 拓展名
	path    string // 原始路径（不含文件名）
	newname string // 新文件名
	archdir string // 归档目录
	id      int    // 文件编辑，默认为0
}

func main() {
	var (
		optPrefix     string // 自定义前缀
		optDir        string // 照片目录
		optVersion    bool   // 打印版本号
		optRenameOnly bool   // 是否只重命名文件，不归档照片
		//optUI         bool   // WEB UI 开关
		optDebug bool // DEBUG模式开关
	)
	ch := make(chan int)

	flag.StringVar(&optPrefix, "prefix", "", "Prefix of file name.")
	flag.StringVar(&optDir, "dir", "./", "Working dir")
	flag.BoolVar(&optRenameOnly, "renameonly", false, "Just rename file.")
	//flag.BoolVar(&optUI, "ui", false, "start the web ui.[PLANNING]")
	flag.BoolVar(&optDebug, "debug", false, "debug mode")
	flag.BoolVar(&optVersion, "version", false, "print version")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	if optDebug {
		log.SetLevel(log.DebugLevel)
		log.Debug("IN DEBUG MODE.")
		//print("XXXXXXX\n")
	}
	if optVersion {
		fmt.Println(VERSION)
		os.Exit(0)
	}
	// --------
	log.Info("VERSION: ", VERSION)
	log.Debug(os.Args)

	// 取得文件列表
	filelist, err := getFileList(optDir)
	if err != nil {
		os.Exit(1)
	}
	index := 0 // 文件记数
	for _, file := range filelist {
		log.Debug("GET FILE = ", file)
		ext := filepath.Ext(file)
		switch ext {
		case ".jpg":
			index++
			log.Debug("GET JPG = ", file)
			// 重命名照片
			go worker(file, optPrefix, optRenameOnly, ch)
		case ".mp4":
			log.Debug(file, "GET A MOV.")
		}
	}
	for i := 0; i < index; i++ {
		<-ch
	}
}

// 取得文件列表
func getFileList(path string) (filelist []string, err error) {
	enteries, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, enterie := range enteries {
		if enterie.IsDir() {
			continue
		}
		file := filepath.Join(path, enterie.Name())
		filelist = append(filelist, file)
	}
	return
}

// 生成新的文件名，返回一个Pic结构体
func genNewFileName(file string, prefix string) (p Pic, err error) {
	p.ext = filepath.Ext(file)
	p.id = 0
	p.name = filepath.Base(file)
	p.path = filepath.Dir(file)
	jpgfile, err := os.Open(file)
	defer jpgfile.Close()
	if err != nil {
		log.Debug(err)
		return
	}

	x, err := exif.Decode(jpgfile)
	if err != nil {
		log.Debug(err)
		return
	}

	datetime, _ := x.DateTime()
	dt2 := strings.Split(datetime.Format("20060102 150405"), " ")
	d := dt2[0]
	t := dt2[1]

	if prefix != "" {
		p.newname = fmt.Sprintf("%s-%s-%s.jpg", prefix, d, t)
	} else {
		p.newname = fmt.Sprintf("%s-%s.jpg", d, t)
	}
	p.archdir = d[:6]
	return p, nil
}

// 主函数，负责重命名照片文件
func worker(path string, prefix string, renameonly bool, ch chan int) {
	var newfullpath string
	p, err := genNewFileName(path, prefix)
	if err != nil {
		log.Debug(err)
		ch <- 1
		return
	}
	if renameonly {
		newfullpath = filepath.Join(p.path, p.newname)
		log.Info("newfullpath = ", newfullpath)
	} else {
		newfullpath = filepath.Join(p.path, p.archdir, p.newname)
		log.Info("newfullpath = ", newfullpath)
		os.MkdirAll(filepath.Join(p.path, p.archdir), 0777)
	}
	// 重命名照片
	err = os.Rename(path, newfullpath)
	if err != nil {
		// 失败直接返回
		log.Infof("- JPG: %s Fatal.\n", path)
		ch <- 1
		return
	}
	log.Infof("+ JPG: %s OK.\n", path)
	ch <- 1
}
