 picar
------------

自家用的照片管理工具，Build with Love。

    注意：
       只在 OSX 和 Windows XP 上测试过。
       得益于 Golang 的特性，理论上也支持 Linux/*BSD/Plan9。

## 特性

1. 依据 EXIF 信息，按拍照日期时间重命名照片。
2. 按**年份月份**归档照片。

    比如： IMG_20140401.jpg => 201504/PREFIX_20150401_111111.jpg

3. Windows 下有一个简单的图形壳，只在 XP 上测试过。 
4. 使用 [Gopm][0] 包管理工具。原因你懂的。
5. [PLANNING]将来会有一个 webui。

## 命令行参数

```
$ picar 命令 [命令选项] [参数]
```

命令:

命令|描述
---|----
`version`       |显示版本号
`rename`        |重命名（及归档）照片文件
`webui`         |启动 webui
`help`, `h`     |显示帮助

```
$ picar rename [命令选项] 目录
```

如果不设置 `目录` 参数，则默认处理当前目录。

`rename` 命令选项：

选项|说明
---|----
`--prefix`         |设置前缀
`--noarchiving`    |是否归档标记
`--debug`          |启动调试模式

### License

WTFPL，详见LICENSE文件。

## 致谢

* 文件处理相关代码参考了[这里][1]。感谢原作者。
* 原生 EXIF 处理库：[rwcarlsen/goexif][2]
* 分级 LOG 库：[Sirupsen/logrus][3]
* [AutoHotKey][4]
* 命令行处理库：[codegangsta/cli][5]

[0]: http://gopm.io/
[1]: http://www.codesnippet.cn/detail/160420132830.html
[2]: https://github.com/rwcarlsen/goexif
[3]: https://github.com/Sirupsen/logrus
[4]: http://ahkscript.org/
[5]: https://github.com/codegangsta/cli
