# picar

自家用照片归档工具，build with love ❤

## 功能

picar 首先按拍照日期时间重命名照片文件，然后按 **年份月份** 归档到不同的目录。
	
重命名文件时会先尝试从照片文件的 EXIF 信息获取数据；如果 EXIF 不存在，则尝试从文件名中获取数据；都失败时会使用当前时间为日期数据。

比如：`IMG_20151106_212111.jpg` 会归档为 `$pwd/201511/20151106_212111.jpg`

Windows 上有一个 [AArdio] 实现图形界面，代码在 [aardio] 目录。

>在 Windows 10 64 位上测试验证 OK；理论上也支持 Linux、macOS。

[AArdio]: http://www.aardio.com/
[aardio]: ./aardio

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

## 安装

```sh
git clone https://github.com/yuekcc/picar.git
cd picar
go install # 需要 Go 1.17 以上版本
```

## 致谢

- 文件处理相关代码参考了 [这里][codesnippet]，感谢原作者
- 原生 EXIF 处理库：[rwcarlsen/goexif]

[rwcarlsen/goexif]: https://github.com/rwcarlsen/goexif
[codesnippet]: http://www.codesnippet.cn/detail/160420132830.html

### License

[MIT](LICENSE)
