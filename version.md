## dotlog版本记录：

#### Version 0.9.9
* Architecture: 调整Logger API定义，日志函数移除Logger返回参数
* Fix: 修正FmtTarget在Error级别时输出两遍内容的问题 
* 2020-01-05 20:00 at ShangHai

#### Version 0.9.8
* Architecture: move xxxxFormat to xxxxF 
* New Feature: add logger.TraceS\DebugS\InfoS\WarnS\ErrorS(content ...interface{}), default will use SprintSpacing format
* Detail:
    - DebugS mean DebugSpacing
    - DebugF mean DebugFormat
* 2019-06-14 11:00

#### Version 0.9.7
* New Feature: add dotlog.SprintSpacing to formats using the default formats 
* 2019-06-14 10:00

#### Version 0.9.6
* Bug fixed: use "Logger.Layout" set info if you config it and use config mode
* 2018-11-02 19:00

#### Version 0.9.5
* 调整：TraceFormat\DebugFormat\InfoFormat\WarnFormat\ErrorFormat新增参数format
* 2018-04-24 09:00

#### Version 0.9.4
* EMailTarget增加MailNickName设置，用于设置发件人友好名称
* 2018-01-19 15:30

#### Version 0.9.3
* 新增FmtTarget，用于向控制台输出日志，同时如果为Error级别消息，同步向StdErr输出
* 模块启动日志增加版本号输出,比如：devfeel.dotlog [0.9.3] InitConfig success
* 新增两个配置模板：ConfigMode_Fmt、ConfigMode_FileFmt，用于简化常规日志配置
* 目前配置模板包含以下：
* ConfigMode_File: 支持Trace\Debug\Info\Warn\Error级别消息，默认输出到文本文件，名称格式：{LogLevel}_{LoggerName}.log
* ConfigMode_Fmt: 支持Trace\Debug\Info\Warn\Error级别消息，默认输出到控制台，如果为Error级别消息，同步向StdErr输出
* ConfigMode_FileFmt: 支持Trace\Debug\Info\Warn\Error级别消息，默认输出到文本文件及控制台，规则参考File与Fmt模板
* ConfigMode_Classics: 支持Trace\Debug\Info\Warn\Error级别消息，默认输出到文本文件，如果为Warn与Error级别，同步输出到邮件
* File输出目录，依赖配置文件中variable:LogDateDir
* EMail输出配置，依赖配置文件中variable:MailServer\variable:ToMail\variable:MailAccount\variable:MailPassword
* 2018-01-19 15:00

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