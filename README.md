 picar
------------

自家用照片管理工具，build with love。

注意：

- 目前只在 Windows 7 64bit 上测试过。
- 得益于 Golang 的特性，理论上也支持 OSX/Linux/*BSD/Plan9。

## 功能

1. 依据 EXIF 信息，按拍照日期时间重命名照片。
2. 按「年份月份」归档照片。

	比如：
	IMG_20151106_212111.jpg => 201504/PREFIX_20150401_111111.jpg

## 使用

```
$ picar 命令 [命令选项] [目录 1] [目录 2] ...
```

命令行参数:

命令 | 描述
--------------------|----------------------
`-prefix [PREFIX]`  | 设置新文件的前缀
`-renameonly`       | 只重命名照片文件，不归档
`-help`, `-h`       | 显示帮助

不带 `[目录]` 参数，就会默认处理当前目录。

### License

WTFPL，详见LICENSE文件。

## 安装

```
$ git clone https://git.coding.net/yuekcc/picar.git $GOPATH/picar
$ go build
```

## 致谢

* 文件处理相关代码参考了[这里][1]，感谢原作者
* 原生 EXIF 处理库：[rwcarlsen/goexif][2]

## TODO

- [x] 完善输出的信息
- [ ] 可以处理视频文件

[0]: http://gopm.io/
[1]: http://www.codesnippet.cn/detail/160420132830.html
[2]: https://github.com/rwcarlsen/goexif