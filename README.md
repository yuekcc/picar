 picar
------------

自家用的照片管理工具，Build with Love。

    注意：
       只在 OSX 和 Windows XP 上测试过。
       得益于 Golang 的特性，理论上也支持 Linux/*BSD/Plan9。
       IE6 的朋友对不住了，web-ui 估计是不支持。

## 功能

1. 依据 EXIF 信息，按拍照日期时间重命名照片。
2. 按「年份月份」归档照片。

    比如： IMG_20140401.jpg => 201504/PREFIX_20150401_111111.jpg

3. 使用 [gopm][0] 包管理工具。原因你可以想象得到。
4. 有一个简单的 web-ui。

## 使用

```
$ picar 命令 [命令选项] [参数]
```

命令:

命令 | 描述
----------------|----------------------
`version`       | 显示版本号
`rename`        | 重命名（及归档）照片文件
`webui`         | 启动 webui
`help`, `h`     | 显示帮助

```
$ picar rename [命令选项] 目录
```

如果不设置 `目录` 参数，则默认处理当前目录。

`rename` 命令选项：

选项 | 说明
-------------------|------------
`--prefix`         | 设置前缀
`--noarchiving`    | 是否归档标记
`--debug`          | 启动调试模式

### License

WTFPL，详见LICENSE文件。

## 安装

安装 picar 前请先安装 [gopm][6] 和 [mkbifs][7]。

```
# cd $PATH_TO_GOPATH
$ git clone https://git.coding.net/yuekcc/picar.git
$ make get
$ make
$ make install
```

## 致谢

* 文件处理相关代码参考了[这里][1]。感谢原作者。
* 原生 EXIF 处理库：[rwcarlsen/goexif][2]
* 分级 LOG 库：[Sirupsen/logrus][3]
* 命令行处理库：[codegangsta/cli][5]

## to-do

- 完善输出的信息
- 可以处理视频文件
- 为 windows 平台加入一个简单的 GUI 壳
- 多国语言如何？
- webui 路径自动提示？

[0]: http://gopm.io/
[1]: http://www.codesnippet.cn/detail/160420132830.html
[2]: https://github.com/rwcarlsen/goexif
[3]: https://github.com/Sirupsen/logrus
[5]: https://github.com/codegangsta/cli
[6]: https://github.com/gpmgo/gopm
[7]: https://coding.net/u/yuekcc/p/mkbifs/git