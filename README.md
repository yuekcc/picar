# picar

自家用照片管理工具，build with love ❤

## 功能

按拍照日期时间重命名照片文件，并按 **年份月份** 归档到不同的目录
	
重命名文件时会先尝试从照片文件的 EXIF 信息获取数据；如果不存在则尝试从文件名中获取数据；都失败时会使用当前的时间日期。

比如：`IMG_20151106_212111.jpg` 会归档为 `$pwd/201511/PREFIX_20150401_111111.jpg`

Windows 上有一个简单的 [AArdio][5] 写的图形界面，代码在 [aardio 目录][6]。

>在 Windows 10 64 位上测试验证 OK；理论上也支持 Linux、macOS。

## 使用

```
$ picar <命令行参数> [目录 1] [目录 2] ...
```

命令行参数：

命令 | 描述
--------------------|----------------------
`-prefix PREFIX`    | 设置新文件的前缀
`-renameonly`       | 只重命名照片文件，不归档
`-videos`			| 处理视频文件
`-help`, `-h`       | 显示帮助

未指定 `[目录]` 参数，处理当前目录。

### License

[MIT](LICENSE)

## 安装

```sh
git clone https://github.com/yuekcc/picar.git
cd picar
go install # 需要 Go 1.17 以上版本
```

## 致谢

* 文件处理相关代码参考了 [这里][1]，感谢原作者
* 原生 EXIF 处理库：[rwcarlsen/goexif][2]


[1]: http://www.codesnippet.cn/detail/160420132830.html
[2]: https://github.com/rwcarlsen/goexif
[3]: https://coding.net/u/yuekcc/p/picar/git/blob/master/LICENSE
[4]: http://tonybai.com/2015/07/31/understand-go15-vendor/
[5]: http://www.aardio.com/
[6]: ./aardio
