## dotlog版本记录：

#### Version 0.9.2
* 配置文件增加Global.ChanSize选项，用于设置File、Mail、Http三类Target的队列通道缓存长度
* 配置方式：<global chansize="1000"></global>
* 当属性值小于等于0，忽略配置值，自动使用默认值默DefaultChanSize=1000
* FileTarget忽略file stat error, fixed for #1 bug: golog.writeTarget os.Stat error
* 2017-11-07 22:00

#### Version 0.9.1
* 输出到File目标时，增加FileMaxSize配置 - 单日志文件最大容量，单位为KB
* 当设置FileMaxSize时，如果发现文件尺寸已达到设置值，则自动将原有文件重命名，后续日志继续往新文件写入
* 重命名规则：假设原文件名为test.log，则自动重命名为 test.log.20060102150405.999999.logbak
* 2017-09-20 18:00

#### Version 0.9
* 初始版本，支持xml配置文件，支持file、http、udp、email四类target
* 2017-08-23 18:00