 picar
------------
自家用照片管理工具，build with love。

注意：

- 目前只在 Windows 7 64 位上测试过。
-  需要 Go 1.5+ 版本编译，因为使用了 [vendor][4] 特性。
- 得益于 Go 的功能，理论上也支持 maxOS、*Linux、*BSD、Plan9。

## 功能

1. 依据 EXIF 信息，按拍照日期时间重命名照片。
2. 按「年份月份」归档照片。

	比如：
	IMG_20151106_212111.jpg => 201511/PREFIX_20150401_111111.jpg

3. Windows 上有一个简单的 [AArdio][5] 写的图形界面，代码在 [aardio 目录][6]。

## 使用

```
$ picar 命令 [命令选项] [目录 1] [目录 2] ...
```

命令行参数:

命令 | 描述
--------------------|----------------------
`-prefix [PREFIX]`  | 设置新文件的前缀
`-noarchiving`      | 只重命名照片文件，不归档
`-videos`			| 处理视频文件
`-help`, `-h`       | 显示帮助

不带 `[目录]` 参数，就会默认处理当前目录。

### License

WTFPL，详见 [LICENSE][3] 文件。

## 安装

```
$ git clone https://git.coding.net/yuekcc/picar.git $GOPATH/src/picar
$ go build
```

## 致谢

* 文件处理相关代码参考了[这里][1]，感谢原作者
* 原生 EXIF 处理库：[rwcarlsen/goexif][2]

## TODO

- [x] 可以处理连拍的照片
- [ ] 可以处理视频文件

[1]: http://www.codesnippet.cn/detail/160420132830.html
[2]: https://github.com/rwcarlsen/goexif
[3]: https://coding.net/u/yuekcc/p/picar/git/blob/master/LICENSE
[4]: http://tonybai.com/2015/07/31/understand-go15-vendor/
[5]: http://www.aardio.com/
[6]: https://coding.net/u/yuekcc/p/picar/git/tree/master/aardio
