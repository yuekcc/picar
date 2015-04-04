picar
------------

自家用的照片管理工具。

本人只在 OSX 和 Windows XP 上测试过。利益于 Go 的特性，理论上也支持 Linux/*BSD。

## 特性

1. 依据 EXIF 信息，按拍照日期时间重命名照片。
2. 按**年份月份**归档照片。

    比如： IMG_20140401.jpg => 201504/PREFIX_20150401_111111.jpg

3. Windows 下有一个简单的图形壳，只在 XP 上测试过。 
4. 使用 [Gopm][0] 包管理工具。原因你懂的。
5. [PLANNING]将来会有一个 webui。

### License

WTFPL，详见LICENSE文件。

## 致谢

* 文件处理相关代码参考了[这里][1]。感谢原作者。
* 原生 EXIF 处理库：[rwcarlsen/goexif][2]
* 分级 LOG 库：[Sirupsen/logrus][3]
* [AutoHotKey][4]

[0]: http://gopm.io/
[1]: http://www.codesnippet.cn/detail/160420132830.html
[2]: https://github.com/rwcarlsen/goexif
[3]: https://github.com/Sirupsen/logrus
[4]: http://ahkscript.org/
